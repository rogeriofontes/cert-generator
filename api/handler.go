package api

import (
	"net/http"
	"strconv"

	"github.com/rogeriofontes/cert-generator/internal/app"
	"github.com/rogeriofontes/cert-generator/internal/utils"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	CertificateService *app.CertificateService
}

// 🔹 Gera certificados para todos os participantes de um evento
func (h *Handler) GenerateCertificatesByEvent(c *gin.Context) {
	// 🔹 Agora usamos a função utilitária para obter a URL base
	baseURL := utils.GetBaseURL(c)

	eventoID, err := strconv.Atoi(c.Param("eventID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do evento inválido"})
		return
	}

	err = h.CertificateService.GenerateCertificatesByEvent(uint(eventoID), baseURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Certificados gerados e enviados para os participantes do evento!"})
}

// 🔹 Gera um certificado para um participante específico
func (h *Handler) GenerateCertificateForUser(c *gin.Context) {
	// 🔹 Agora usamos a função utilitária para obter a URL base
	baseURL := utils.GetBaseURL(c)

	usuarioID, err := strconv.Atoi(c.Param("usuarioID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do participante inválido"})
		return
	}

	err = h.CertificateService.GenerateCertificateForUser(uint(usuarioID), baseURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Certificado gerado e enviado para o participante!"})
}

// 🔹 Gera certificados para todos os participantes pendentes
func (h *Handler) GeneratePendingCertificates(c *gin.Context) {
	// 🔹 Agora usamos a função utilitária para obter a URL base
	baseURL := utils.GetBaseURL(c)

	err := h.CertificateService.GenerateAllPendingCertificates(baseURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todos os certificados pendentes foram gerados e enviados!"})
}
