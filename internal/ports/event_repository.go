package ports

import "github.com/rogeriofontes/cert-generator/internal/domain"

type EventRepository interface {
	Save(event *domain.Event) error
	FindByID(id uint) (*domain.Event, error)
	FindAll() ([]domain.Event, error)
}
