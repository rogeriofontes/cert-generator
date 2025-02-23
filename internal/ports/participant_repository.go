package ports

import "github.com/rogeriofontes/cert-generator/internal/domain"

type ParticipantRepository interface {
	Save(participant *domain.Participant) error
	FindAll() ([]domain.Participant, error)
	FindAllPending() ([]domain.Participant, error)
	FindByEvent(eventID uint) ([]domain.Participant, error)
	FindByID(id uint) (*domain.Participant, error)
	FindByIDWithEvent(id uint) (*domain.Participant, error) // 🔹 Novo método para buscar participante com evento
	Update(participant *domain.Participant) error
	UpdateParticipantCertificateId(participantID uint, certificateId string) error
	FindByCertificateId(code string) (*domain.Participant, error)
}
