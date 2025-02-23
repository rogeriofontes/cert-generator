package db

import (
	"errors"

	"github.com/rogeriofontes/cert-generator/internal/domain"
	"gorm.io/gorm"
)

// UserRepository define as operações para o usuário
type UserRepo struct {
	DB *gorm.DB
}

// CreateUser cria um novo usuário
func (r *UserRepo) CreateUser(user *domain.User) error {
	return r.DB.Create(user).Error
}

// GetUserByEmail busca um usuário pelo e-mail
func (r *UserRepo) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// GetUserByID busca um usuário pelo ID
func (r *UserRepo) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.DB.First(&user, id).Error
	return &user, err
}

// GetAllUsers retorna todos os usuários
func (r *UserRepo) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	err := r.DB.Find(&users).Error
	return users, err
}

// DeleteUser remove um usuário pelo ID
func (r *UserRepo) DeleteUser(id uint) error {
	return r.DB.Delete(&domain.User{}, id).Error
}
