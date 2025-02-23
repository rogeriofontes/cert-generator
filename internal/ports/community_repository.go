package ports

import "github.com/rogeriofontes/cert-generator/internal/domain"

type CommunityRepository interface {
	Save(comunity *domain.Community) error
	FindByID(id uint) (*domain.Community, error)
	FindAll() ([]domain.Community, error)
}
