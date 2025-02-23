package ports

import "github.com/rogeriofontes/cert-generator/internal/domain"

type EmailService interface {
	SendEmail(participant domain.Participant, certificadoPath string) error
}
