package app

import (
	"errors"

	"github.com/rogeriofontes/cert-generator/internal/auth"
	"github.com/rogeriofontes/cert-generator/internal/domain"
	"github.com/rogeriofontes/cert-generator/internal/ports"
	"golang.org/x/crypto/bcrypt"
)

// UserService contém a lógica de negócios para usuários
type UserService struct {
	UserRepo ports.UserRepository
}

// RegisterUser registra um novo usuário com senha criptografada
func (s *UserService) RegisterUser(user *domain.User) error {
	// Verificar se o e-mail já está cadastrado
	existingUser, _ := s.UserRepo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("e-mail já cadastrado")
	}

	// Criptografar senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return s.UserRepo.CreateUser(user)
}

// AuthenticateUser verifica as credenciais e retorna um token JWT
func (s *UserService) AuthenticateUser(email, password string) (string, error) {
	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("usuário ou senha inválidos")
	}

	// Verificar senha
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("usuário ou senha inválidos")
	}

	// Gerar token JWT
	return auth.GenerateJWT(user.Email, user.Role)
}

// GetAllUsers retorna todos os usuários
func (s *UserService) GetAllUsers() ([]domain.User, error) {
	return s.UserRepo.GetAllUsers()
}
