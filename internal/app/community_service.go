package app

import (
	"fmt"

	"github.com/rogeriofontes/cert-generator/internal/domain"
	"github.com/rogeriofontes/cert-generator/internal/ports"
)

type CommunityService struct {
	CommunityRepo ports.CommunityRepository
}

// Criar um novo evento
func (s *CommunityService) CreateCommunity(community *domain.Community) error {
	if community.Name == "" {
		return fmt.Errorf("o nome e a data do evento são obrigatórios")
	}
	return s.CommunityRepo.Save(community)
}

// Buscar todos os eventos
func (s *CommunityService) GetAllCommunits() ([]domain.Community, error) {
	return s.CommunityRepo.FindAll()
}
