package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rogeriofontes/cert-generator/internal/auth"
)

// JWTMiddleware verifica se o usuário está autenticado
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obter token do cabeçalho Authorization: Bearer <token>
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token ausente"})
			c.Abort()
			return
		}

		// Remover "Bearer " e obter apenas o token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato do token inválido"})
			c.Abort()
			return
		}

		// Validar o token JWT
		claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Adicionar claims no contexto para uso nas rotas protegidas
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}
