package db

import (
	"github.com/rogeriofontes/cert-generator/internal/domain"

	"gorm.io/gorm"
)

// EventoRepo representa o repositório para eventos
type EventRepo struct {
	DB *gorm.DB
}

// Salvar um evento
func (r *EventRepo) Save(event *domain.Event) error {
	return r.DB.Create(event).Error
}

// Buscar evento pelo ID
func (r *EventRepo) FindByID(id uint) (*domain.Event, error) {
	var event domain.Event
	err := r.DB.First(&event, id).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

// Buscar todos os eventos
func (r *EventRepo) FindAll() ([]domain.Event, error) {
	var events []domain.Event
	err := r.DB.Find(&events).Error
	return events, err
}
