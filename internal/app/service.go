package app

import (
	"fmt"
	"log"

	"github.com/rogeriofontes/cert-generator/internal/ports"
)

type CertificateService struct {
	EventRepo       ports.EventRepository
	ParticipantRepo ports.ParticipantRepository
	PdfGen          ports.PDFGenerator
	EmailSvc        ports.EmailService
}

// 📌 Gerar certificados para todos os participantes de um evento
func (s *CertificateService) GenerateCertificatesByEvent(eventID uint, baseURL string) error {
	participants, err := s.ParticipantRepo.FindByEvent(eventID)
	if err != nil {
		return fmt.Errorf("erro ao buscar participantes: %v", err)
	}
	if len(participants) == 0 {
		return fmt.Errorf("nenhum participante encontrado para o evento %d", eventID)
	}

	for _, participant := range participants {
		if participant.Status == "pendente" {
			filePath, err := s.PdfGen.GenerateCertificate(&participant, baseURL)
			if err != nil {
				log.Printf("Erro ao gerar certificado para %s: %v", participant.Name, err)
				continue
			}

			err = s.EmailSvc.SendEmail(participant, filePath)
			if err != nil {
				log.Printf("Erro ao enviar e-mail para %s: %v", participant.Email, err)
				continue
			}

			participant.Status = "gerado"
			participant.Certificate = filePath
			s.ParticipantRepo.Update(&participant)
		}
	}

	log.Printf("Certificados gerados e enviados para o evento %d", eventID)
	return nil
}

// 🔹 Gera certificados para todos os participantes pendentes
func (s *CertificateService) GenerateAllPendingCertificates(baseURL string) error {
	// 1️⃣ Busca todos os participantes com status "pendente"
	participants, err := s.ParticipantRepo.FindAllPending()
	if err != nil {
		return fmt.Errorf("erro ao buscar participantes pendentes: %v", err)
	}

	// Se não houver participantes pendentes, retorna um erro amigável
	if len(participants) == 0 {
		return fmt.Errorf("nenhum certificado pendente encontrado")
	}

	// 2️⃣ Processa cada participante pendente
	for _, participant := range participants {
		log.Printf("🔹 Gerando certificado para: %s", participant.Name)

		// 3️⃣ Gera o certificado em PDF
		filePath, err := s.PdfGen.GenerateCertificate(&participant, baseURL)
		if err != nil {
			log.Printf("❌ Erro ao gerar certificado para %s: %v", participant.Name, err)
			continue // Pula para o próximo participante se houver erro
		}

		// 4️⃣ Envia o certificado por e-mail
		err = s.EmailSvc.SendEmail(participant, filePath)
		if err != nil {
			log.Printf("❌ Erro ao enviar e-mail para %s: %v", participant.Email, err)
			continue // Pula para o próximo participante se houver erro
		}

		// 5️⃣ Atualiza o status do participante para "gerado"
		participant.Status = "gerado"
		participant.Certificate = filePath
		err = s.ParticipantRepo.Update(&participant)
		if err != nil {
			log.Printf("❌ Erro ao atualizar status do participante %s: %v", participant.Name, err)
		} else {
			log.Printf("✅ Certificado gerado e enviado para %s", participant.Name)
		}
	}

	log.Println("🎉 Todos os certificados pendentes foram gerados e enviados com sucesso!")
	return nil
}

// 📌 Gerar um certificado para um participante específico
func (s *CertificateService) GenerateCertificateForUser(usuarioID uint, baseURL string) error {
	participante, err := s.ParticipantRepo.FindByID(usuarioID)
	if err != nil {
		return fmt.Errorf("participante não encontrado: %v", err)
	}

	if participante.Status == "pendente" {
		filePath, err := s.PdfGen.GenerateCertificate(participante, baseURL)
		if err != nil {
			return fmt.Errorf("erro ao gerar certificado: %v", err)
		}

		err = s.EmailSvc.SendEmail(*participante, filePath)
		if err != nil {
			return fmt.Errorf("erro ao enviar e-mail: %v", err)
		}

		participante.Status = "gerado"
		participante.Certificate = filePath
		s.ParticipantRepo.Update(participante)
	}

	log.Printf("Certificado gerado e enviado para o usuário %d", usuarioID)
	return nil
}
