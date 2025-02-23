package app

import (
	"fmt"

	"github.com/rogeriofontes/cert-generator/internal/domain"
	"github.com/rogeriofontes/cert-generator/internal/ports"
)

type ParticipantService struct {
	ParticipantRepo ports.ParticipantRepository
}

// Criar um novo participante em um evento
func (s *ParticipantService) CreateParticipant(participant *domain.Participant) error {
	if participant.Name == "" || participant.Email == "" || participant.EventID == 0 {
		return fmt.Errorf("nome, email e evento_id são obrigatórios")
	}
	return s.ParticipantRepo.Save(participant)
}

func (s *ParticipantService) GetParticipants() ([]domain.Participant, error) {
	return s.ParticipantRepo.FindAll()
}

// Buscar todos os participantes de um evento específico
func (s *ParticipantService) GetParticipantsByEvent(eventID uint) ([]domain.Participant, error) {
	return s.ParticipantRepo.FindByEvent(eventID)
}

// Buscar participante com evento
func (s *ParticipantService) GetParticipantByEvent(id uint) (*domain.Participant, error) {
	participant, err := s.ParticipantRepo.FindByIDWithEvent(id)
	if err != nil {
		return nil, fmt.Errorf("participante não encontrado: %v", err)
	}
	return participant, nil
}

// Atualiza o CertificateId do participante no banco de dados
func (s *ParticipantService) UpdateParticipantCertificateId(participantID uint, certificateId string) error {
	return s.ParticipantRepo.UpdateParticipantCertificateId(participantID, certificateId)
}

// FindByCertificateId busca um participante pelo código UUID do certificado
func (s *ParticipantService) FindByCertificateId(code string) (*domain.Participant, error) {
	return s.ParticipantRepo.FindByCertificateId(code)
}
