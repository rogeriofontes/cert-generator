package email

import (
	"fmt"

	"github.com/rogeriofontes/cert-generator/internal/domain"

	"gopkg.in/gomail.v2"
)

type EmailServiceImpl struct {
	SMTPUser string
	SMTPPass string
}

func (e *EmailServiceImpl) SendEmail(participante domain.Participant, certificadoPath string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.SMTPUser)
	m.SetHeader("To", participante.Email)
	m.SetHeader("Subject", "Certificado de Participação")
	m.SetBody("text/plain", fmt.Sprintf("Olá %s, segue seu certificado.", participante.Name))
	m.Attach(certificadoPath)

	d := gomail.NewDialer("smtp.gmail.com", 587, e.SMTPUser, e.SMTPPass)
	return d.DialAndSend(m)
}
