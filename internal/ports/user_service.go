package ports

import "github.com/rogeriofontes/cert-generator/internal/domain"

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByEmail(email string) (*domain.User, error)
	GetUserByID(id uint) (*domain.User, error)
	GetAllUsers() ([]domain.User, error)
	DeleteUser(id uint) error
}
