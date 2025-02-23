package api

import (
	"net/http"

	"github.com/rogeriofontes/cert-generator/internal/app"
	"github.com/rogeriofontes/cert-generator/internal/domain"

	"github.com/gin-gonic/gin"
)

// EventHandler é o handler específico para eventos
// EventHandler é o handler específico para eventos
// @Summary Handler para eventos
// @Description Handler para eventos
// @Tags Eventos
// @Accept json
// @Produce json
// @Router /events [get]
// @Router /events [post]
// @Router /events/{id} [get]
type EventHandler struct {
	EventService *app.EventService
}

// 🔹 Criar um novo evento
// @Summary Cadastrar um novo evento
// @Description Cadastrar um novo evento
// @Tags Eventos
// @Accept json
// @Produce json
// @Param event body domain.Event true "Evento a ser cadastrado"
// @Success 201 {object} map[string]string "Evento cadastrado com sucesso"
// @Failure 400 {object} map[string]string "Dados inválidos"
// @Failure 500 {object} map[string]string "Erro ao cadastrar evento"
// @Router /events [post]
func (h *EventHandler) CreateEvent(c *gin.Context) {
	// Garantir que a requisição está usando UTF-8
	c.Header("Content-Type", "application/json; charset=utf-8")

	var event domain.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	err := h.EventService.CreateEvent(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Evento criado com sucesso!"})
}

// 🔹 Buscar todos os eventos
// @Summary Buscar todos os eventos
// @Description Buscar todos os eventos
// @Tags Eventos
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Event
// @Router /events [get]
func (h *EventHandler) GetEvents(c *gin.Context) {
	events, err := h.EventService.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar eventos"})
		return
	}

	c.JSON(http.StatusOK, events)
}
