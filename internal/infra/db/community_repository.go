package db

import (
	"github.com/rogeriofontes/cert-generator/internal/domain"

	"gorm.io/gorm"
)

// ComunidadeRepo representa o repositório para eventos
type CommunityRepo struct {
	DB *gorm.DB
}

// Salvar uma comunidade
func (r *CommunityRepo) Save(community *domain.Community) error {
	return r.DB.Create(community).Error
}

// Buscar comunidade pelo ID
func (r *CommunityRepo) FindByID(id uint) (*domain.Community, error) {
	var community domain.Community
	err := r.DB.First(&community, id).Error
	if err != nil {
		return nil, err
	}
	return &community, nil
}

// Buscar todas as comunidade
func (r *CommunityRepo) FindAll() ([]domain.Community, error) {
	var communities []domain.Community
	err := r.DB.Find(&communities).Error
	return communities, err
}
