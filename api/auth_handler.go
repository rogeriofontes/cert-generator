package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rogeriofontes/cert-generator/internal/auth"
)

// LoginRequest representa a estrutura do login
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Mock de usuários (substitua por um banco de dados)
var users = map[string]string{
	"admin": "123456", // username: senha
	"user":  "password",
}

// Login autentica o usuário e retorna um token JWT
// @Summary Faz login na API
// @Description Gera um token JWT após login bem-sucedido
// @Tags Autenticação
// @Accept json
// @Produce json
// @Param login body LoginRequest true "Dados de login"
// @Success 200 {object} map[string]string "Token gerado com sucesso"
// @Failure 400 {object} map[string]string "Erro de autenticação"
// @Router /login [post]
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Verificar se usuário e senha estão corretos
	password, exists := users[req.Username]
	if !exists || password != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha incorretos"})
		return
	}

	// Definir o "role" do usuário (admin ou padrão)
	role := "user"
	if req.Username == "admin" {
		role = "admin"
	}

	// Gerar token JWT
	token, err := auth.GenerateJWT(req.Username, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
