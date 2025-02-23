package db

import (
	"github.com/rogeriofontes/cert-generator/internal/domain"

	"gorm.io/gorm"
)

// ParticipanteRepo representa o repositório para participantes
type ParticipantRepo struct {
	DB *gorm.DB
}

// Salvar um participante
func (r *ParticipantRepo) Save(participant *domain.Participant) error {
	return r.DB.Create(participant).Error
}

// Buscar todos os participantes pendentes
func (r *ParticipantRepo) FindAll() ([]domain.Participant, error) {
	var participants []domain.Participant
	err := r.DB.Find(&participants).Error
	return participants, err
}

// Buscar todos os participantes pendentes
func (r *ParticipantRepo) FindAllPending() ([]domain.Participant, error) {
	var participants []domain.Participant
	err := r.DB.Where("status = ?", "pendente").Find(&participants).Error
	return participants, err
}

// Buscar participantes de um evento específico
func (r *ParticipantRepo) FindByEvent(eventID uint) ([]domain.Participant, error) {
	var participants []domain.Participant
	err := r.DB.Preload("Event").Where("event_id = ?", eventID).Find(&participants).Error
	return participants, err
}

// Buscar um participante pelo ID
func (r *ParticipantRepo) FindByID(id uint) (*domain.Participant, error) {
	var participant domain.Participant
	err := r.DB.First(&participant, id).Error
	if err != nil {
		return nil, err
	}
	return &participant, nil
}

// Buscar um participante e carregar os dados do evento
func (r *ParticipantRepo) FindByIDWithEvent(id uint) (*domain.Participant, error) {
	var participant domain.Participant
	err := r.DB.Preload("Event").First(&participant, id).Error
	if err != nil {
		return nil, err
	}
	return &participant, nil
}

// Atualizar um participante (ex: mudar status para "gerado" e adicionar certificado)
func (r *ParticipantRepo) Update(participant *domain.Participant) error {
	return r.DB.Save(participant).Error
}

// Atualiza o CertificateId do participante
func (r *ParticipantRepo) UpdateParticipantCertificateId(participantID uint, certificateId string) error {
	return r.DB.Model(&domain.Participant{}).
		Where("id = ?", participantID).
		Update("certificate_id", certificateId).Error
}

// FindByCertificateId busca um participante pelo código do certificado
func (r *ParticipantRepo) FindByCertificateId(code string) (*domain.Participant, error) {
	var participant domain.Participant
	err := r.DB.Preload("Event").Where("certificate_id = ?", code).First(&participant).Error
	if err != nil {
		return nil, err
	}
	return &participant, nil
}
