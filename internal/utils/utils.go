package utils

import "github.com/gin-gonic/gin"

// GetBaseURL retorna a URL base da aplicação a partir do contexto do Gin
func GetBaseURL(c *gin.Context) string {
	host := c.Request.Host // Exemplo: "localhost:9393"
	scheme := "http"       // Default para HTTP

	if c.Request.TLS != nil { // Verifica se é HTTPS
		scheme = "https"
	}

	return scheme + "://" + host
}
