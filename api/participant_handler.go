package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rogeriofontes/cert-generator/internal/app"
	"github.com/rogeriofontes/cert-generator/internal/domain"

	"github.com/gin-gonic/gin"
)

// ParticipantHandler é o handler específico para participantes
// ParticipantHandler é o handler específico para participantes
// @Summary Handler para participantes
// @Description Handler para participantes
// @Tags Participantes
// @Accept json
// @Produce json
// @Router /participants [get]
// @Router /participants [post]
// @Router /events/{eventID}/participants [get]
// @Router /participants/{id} [get]
// @Router /participants/validate [get]
type ParticipantHandler struct {
	ParticipantService *app.ParticipantService
}

// 🔹 Criar um novo participante em um evento
// @Summary Cadastrar um novo participante
// @Description Cadastrar um novo participante em um evento
// @Tags Participantes
// @Accept json
// @Produce json
// @Param participant body domain.Participant true "Participante a ser cadastrado"
// @Success 201 {object} map[string]string "Participante cadastrado com sucesso"
// @Failure 400 {object} map[string]string "Dados inválidos"
// @Failure 500 {object} map[string]string "Erro ao cadastrar participante"
// @Router /participants [post]
func (h *ParticipantHandler) CreateParticipant(c *gin.Context) {
	// Garantir que a requisição está usando UTF-8
	c.Header("Content-Type", "application/json; charset=utf-8")

	var participant domain.Participant

	// 🔹 Apenas faz o binding sem consumir o JSON antes
	if err := c.ShouldBindJSON(&participant); err != nil {
		fmt.Println("❌ Erro ao fazer binding:", err) // Log do erro
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos", "detalhe": err.Error()})
		return
	}

	// 📌 Log do participante antes de salvar
	fmt.Printf("✅ Participante validado: %+v\n", participant)

	err := h.ParticipantService.CreateParticipant(&participant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Participante cadastrado com sucesso!"})
}

// GetParticipants retorna a lista de participantes
// @Summary Buscar todos os participantes
// @Description Buscar todos os participantes cadastrados
// @Tags Participantes
// @Accept json
// @Produce json
// @Success 200 {array} domain.Participant "Participantes encontrados"
// @Failure 500 {object} map[string]string "Erro ao buscar participantes"
// @Router /participants [get]
func (h *ParticipantHandler) GetParticipants(c *gin.Context) {
	participants, err := h.ParticipantService.ParticipantRepo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar participantes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"participants": participants})
}

// 🔹 Buscar todos os participantes de um evento específico
// @Summary("Buscar todos os participantes de um evento")
// @Description("Buscar todos os participantes de um evento pelo ID")
// @Tags("Participantes")
// @Accept(json)
// @Produce(json)
// @Param eventID path int true "ID do evento"
// @Success 200 {array} domain.Participant "Participantes encontrados"
// @Failure 400 {object} map[string]string "ID do evento inválido"
// @Failure 500 {object} map[string]string "Erro ao buscar participantes"
// @Router /events/{eventID}/participants [get]
func (h *ParticipantHandler) GetParticipantesByEvent(c *gin.Context) {
	eventID, err := strconv.Atoi(c.Param("eventID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do evento inválido"})
		return
	}

	participants, err := h.ParticipantService.GetParticipantsByEvent(uint(eventID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar participantes"})
		return
	}

	c.JSON(http.StatusOK, participants)
}

// 🔹 Buscar um participante e seu evento
// @Summary("Buscar um participante e seu evento")
// @Description("Buscar um participante e seu evento pelo ID")
// @Tags("Participantes")
// @Accept(json)
// @Produce(json)
// @Param id path int true "ID do participante"
// @Success 200 {object} domain.Participant "Participante encontrado"
// @Failure 400 {object} map[string]string "ID inválido"
// @Failure 404 {object} map[string]string "Participante não encontrado"
// @Router /participants/{id} [get]
func (h *ParticipantHandler) GetParticipantByEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	participant, err := h.ParticipantService.GetParticipantByEvent(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Participante não encontrado"})
		return
	}

	c.JSON(http.StatusOK, participant)
}

// ValidateParticipantCertificate verifica se o código do certificado é válido
// ValidateParticipantCertificate verifica se o código do certificado é válido
// @Summary Valida um certificado
// @Description Valida um certificado pelo código UUID gerado
// @Tags Certificados
// @Accept json
// @Produce json
// @Param code query string true "Código do certificado"
// @Success 200 {object} map[string]interface{} "Certificado válido"
// @Failure 400 {object} map[string]string "Código ausente"
// @Failure 404 {object} map[string]string "Certificado não encontrado"
// @Router /participants/validate [get]
func (h *ParticipantHandler) ValidateParticipantCertificate(c *gin.Context) {
	// 🔹 Capturar o código UUID da query string
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Código de validação ausente"})
		return
	}

	// 🔹 Buscar o participante pelo código
	participant, err := h.ParticipantService.FindByCertificateId(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Certificado não encontrado"})
		return
	}

	// 🔹 Retornar os detalhes do participante e evento
	c.JSON(http.StatusOK, gin.H{
		"message":        "Certificado válido",
		"participant":    participant.Name,
		"email":          participant.Email,
		"event":          participant.Event.Name,
		"total_hours":    participant.Event.TotalHours,
		"certificate_id": participant.CertificateId,
	})
}
