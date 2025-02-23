package ports

import "github.com/rogeriofontes/cert-generator/internal/domain"

type PDFGenerator interface {
	GenerateCertificate(participant *domain.Participant, baseURL string) (string, error)
}
