package app

import (
	"fmt"

	"github.com/rogeriofontes/cert-generator/internal/domain"
	"github.com/rogeriofontes/cert-generator/internal/ports"
)

type EventService struct {
	EventRepo ports.EventRepository
}

// Criar um novo evento
func (s *EventService) CreateEvent(evento *domain.Event) error {
	if evento.Name == "" || evento.Date == "" {
		return fmt.Errorf("o nome e a data do evento são obrigatórios")
	}
	return s.EventRepo.Save(evento)
}

// Buscar todos os eventos
func (s *EventService) GetAllEvents() ([]domain.Event, error) {
	return s.EventRepo.FindAll()
}
