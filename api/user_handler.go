package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogeriofontes/cert-generator/internal/app"
	"github.com/rogeriofontes/cert-generator/internal/domain"
)

// UserHandler gerencia os usuários
type UserHandler struct {
	UserService *app.UserService
}

// Register cria um novo usuário
// @Summary Registra um usuário
// @Tags Usuários
// @Accept json
// @Produce json
// @Param user body domain.User true "Dados do usuário"
// @Success 201 {object} map[string]string "Usuário registrado com sucesso"
// @Failure 400 {object} map[string]string "Erro ao registrar usuário"
// @Router /users/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	err := h.UserService.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário registrado com sucesso"})
}

// Login faz autenticação e retorna um JWT
// @Summary Faz login
// @Tags Usuários
// @Accept json
// @Produce json
// @Param credentials body map[string]string true "Credenciais de login"
// @Success 200 {object} map[string]string "Token gerado"
// @Failure 401 {object} map[string]string "Usuário ou senha inválidos"
// @Router /users/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var credentials map[string]string
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	email, password := credentials["email"], credentials["password"]
	token, err := h.UserService.AuthenticateUser(email, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
