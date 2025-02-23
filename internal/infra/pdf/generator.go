package pdf

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
	"github.com/rogeriofontes/cert-generator/internal/domain"
	"github.com/rogeriofontes/cert-generator/internal/infra/db"
	"github.com/skip2/go-qrcode"
)

// PDFService gerencia a geração de certificados em PDF
type PDFService struct {
	BackgroundPath  string
	OutputDir       string
	ParticipantRepo *db.ParticipantRepo
}

// GenerateCertificate cria um certificado para um participante
func (s *PDFService) GenerateCertificate(participant *domain.Participant, baseURL string) (string, error) {
	// Garantir que a pasta fonts/ está na raiz do projeto
	projectRoot, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("erro ao obter diretório do projeto: %v", err)
	}
	fontsPath := filepath.Join(projectRoot, "fonts")          // Define a pasta raiz correta
	qrCodePath := filepath.Join(projectRoot, "qrcode")        // Pasta para QR Codes
	signaturePath := filepath.Join(projectRoot, "signatures") // Pasta para Assinaturas

	// Criar diretórios se não existirem
	dirs := []string{s.OutputDir, fontsPath}
	for _, dir := range dirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, os.ModePerm)
			log.Printf("📁 Diretório criado: %s", dir)
		}
	}

	// Gerar um UUID para cada certificado
	uuidCode := uuid.New().String()
	// 🔹 Gerar a URL dinâmica com a base obtida do contexto
	validationURL := fmt.Sprintf("%s/participants/validate?code=%s", baseURL, uuidCode)

	// Gerar QR Code
	qrCodeFile := filepath.Join(qrCodePath, fmt.Sprintf("%s.png", uuidCode))
	err = qrcode.WriteFile(validationURL, qrcode.Medium, 256, qrCodeFile)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar QR Code: %v", err)
	}
	log.Printf("✅ QR Code gerado: %s", qrCodeFile)

	// Caminho absoluto da imagem de fundo
	absBackgroundPath, err := filepath.Abs(s.BackgroundPath)
	if err != nil {
		return "", fmt.Errorf("erro ao obter o caminho absoluto da imagem: %v", err)
	}

	// Criar PDF
	pdf := gofpdf.New("L", "mm", "A4", "")

	// Definir caminho absoluto da fonte Arial
	fontPath := filepath.Join(fontsPath, "arial.ttf")
	absFontPath, err := filepath.Abs(fontPath)
	if err != nil {
		return "", fmt.Errorf("erro ao obter caminho absoluto da fonte: %v", err)
	}

	// Verificar se a fonte Arial existe
	if _, err := os.Stat(absFontPath); os.IsNotExist(err) {
		log.Printf("⚠️ Fonte Arial não encontrada em: %s", absFontPath)
		log.Println("Usando fonte alternativa DejaVuSans.ttf")

		fontPath = filepath.Join(fontsPath, "DejaVuSans.ttf")
		absFontPath, err = filepath.Abs(fontPath)
		if err != nil {
			return "", fmt.Errorf("erro ao obter caminho absoluto da fonte alternativa: %v", err)
		}

		if _, err := os.Stat(absFontPath); os.IsNotExist(err) {
			return "", fmt.Errorf("❌ Nenhuma fonte disponível. Certifique-se de que arial.ttf ou DejaVuSans.ttf estão na pasta fonts/")
		}
	}

	// Log do caminho correto da fonte
	log.Printf("✅ Fonte carregada: %s", absFontPath)

	// Adicionar fonte UTF-8 personalizada (normal e negrito)
	pdf.AddUTF8Font("CustomFont", "", absFontPath)      // Fonte normal
	pdf.AddUTF8Font("CustomFont-Bold", "", absFontPath) // Fonte "B" (negrito)

	// Adicionar página ao PDF
	pdf.AddPage()

	// Adicionar imagem de fundo
	pdf.Image(absBackgroundPath, 0, 0, 297, 210, false, "", 0, "")

	// Definir título
	pdf.SetFont("CustomFont-Bold", "", 28) // Negrito corrigido
	pdf.SetTextColor(0, 0, 0)
	pdf.SetXY(50, 50)
	pdf.CellFormat(200, 10, "Certificado de Participação", "", 0, "C", false, 0, "")

	// Adicionar mensagem personalizada
	texto := "Certificamos que,"
	pdf.SetFont("CustomFont", "", 18)
	pdf.SetXY(30, 80)
	pdf.MultiCell(240, 10, texto, "", "C", false)

	// Nome do Participante em Destaque (Grande e Negrito)
	pdf.SetFont("CustomFont-Bold", "", 30) // Fonte maior e negrito
	pdf.SetXY(30, 90)
	pdf.MultiCell(240, 15, participant.Name, "", "C", false) // Nome separado e maior

	// Adicionar o restante do texto
	texto2 := "participou do evento: "
	pdf.SetFont("CustomFont", "", 18)
	pdf.SetXY(30, 105)
	pdf.MultiCell(240, 10, texto2, "", "C", false)

	// Nome do Evento em Destaque
	pdf.SetFont("CustomFont-Bold", "", 24) // Fonte média e negrito
	pdf.SetXY(30, 115)
	pdf.MultiCell(240, 12, participant.Event.Name, "", "C", false)

	// Adicionar duração e compromisso
	texto3 := "Este com duração de " + fmt.Sprint(participant.Event.TotalHours) + " horas, demonstrando comprometimento com seu desenvolvimento profissional."
	pdf.SetFont("CustomFont", "", 18)
	pdf.SetXY(30, 128)
	pdf.MultiCell(240, 10, texto3, "", "C", false)

	// Adicionar duração e compromisso
	texto4 := "ID de Validação: " + fmt.Sprint(participant.CertificateId)
	pdf.SetFont("CustomFont", "", 8) // Fonte menor para o ID
	pdf.SetXY(20, 179)               // 🔹 10mm da borda esquerda, 200mm de altura (perto do rodapé)
	pdf.CellFormat(0, 10, texto4, "", 0, "L", false, 0, "")

	// Adicionar QR Code no canto direito inferior
	absQRCodePath, err := filepath.Abs(qrCodeFile)
	if err == nil {
		pdf.Image(absQRCodePath, 240, 158, 30, 30, false, "", 0, "")
	}

	// 🔹 Adicionar Assinatura Centralizada com Imagem
	signatureFile := filepath.Join(signaturePath, "assinatura.png") // Nome do arquivo da assinatura
	absSignaturePath, err := filepath.Abs(signatureFile)
	if err == nil {
		// Adiciona a imagem da assinatura no centro da página
		pdf.Image(absSignaturePath, 130, 158, 50, 20, false, "", 0, "")
	}

	// 🔹 Nome do Signatário Abaixo da Assinatura
	pdf.SetXY(130, 178)
	pdf.SetFont("CustomFont-Bold", "", 14)
	pdf.CellFormat(50, 10, "Coordenador do Evento", "", 0, "C", false, 0, "")

	// Criar nome do arquivo, removendo caracteres especiais
	safeName := strings.ReplaceAll(participant.Name, " ", "_")
	fileName := fmt.Sprintf("%s/certificado_%s.pdf", s.OutputDir, safeName)
	absFileName, err := filepath.Abs(fileName)
	if err != nil {
		return "", fmt.Errorf("erro ao obter caminho absoluto do arquivo: %v", err)
	}

	// Salvar PDF
	err = pdf.OutputFileAndClose(absFileName)
	if err != nil {
		return "", fmt.Errorf("erro ao salvar certificado: %v", err)
	}

	// Atualizar o CertificateId do participante no banco de dados
	participant.CertificateId = uuidCode
	err = s.UpdateParticipantCertificateId(participant.ID, uuidCode)
	if err != nil {
		log.Printf("❌ Erro ao atualizar o CertificateId do participante %d: %v", participant.ID, err)
		return "", fmt.Errorf("erro ao atualizar CertificateId: %v", err)
	}

	log.Printf("✅ CertificateId atualizado para o participante %d: %s", participant.ID, uuidCode)

	log.Printf("✅ Certificado gerado: %s", absFileName)
	return absFileName, nil
}

// UpdateParticipantCertificateId atualiza o CertificateId no banco de dados
func (s *PDFService) UpdateParticipantCertificateId(participantID uint, certificateId string) error {
	if s.ParticipantRepo == nil { // 🔹 Verifica se o repositório está inicializado
		log.Fatal("❌ ParticipantRepo está NIL dentro do PDFService")
	}
	return s.ParticipantRepo.UpdateParticipantCertificateId(participantID, certificateId)
}
