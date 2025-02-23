package api

import (
	"net/http"

	"github.com/rogeriofontes/cert-generator/internal/app"
	"github.com/rogeriofontes/cert-generator/internal/domain"

	"github.com/gin-gonic/gin"
)

// ComunidadeHandler é o handler específico para Comunidadeos
// Ele é responsável por receber as requisições HTTP e chamar os métodos apropriados do serviço ComunidadeService
// para executar a lógica de negócios e retornar a resposta apropriada ao cliente.
// O handler é responsável por serializar e desserializar os dados e garantir que os dados estejam corretos.
// @Summary Handler para Comunidadeos
// @Description Handler para Comunidadeos
// @Tags Comunidade
// @Accept json
// @Produce json
// @Router /community [get]
// @Router /community [post]
type CommunityHandler struct {
	CommunityService *app.CommunityService
}

// 🔹 Criar um novo Comunidadeo
// @Summary Criar um novo Comunidadeo
// @Description	ption Criar um novo Comunidadeo
// @Tags Comunidade
// @Accept json
// @Produce json
// @Param community body domain.Community true "Comunidade a ser criada"
// @Success 201 {string} string "Comunidade criada com sucesso!"
// @Failure 400 {string} string "Dados inválidos"
// @Failure 500 {string} string "Erro ao criar Comunidadeo"
// @Router /communities [post]
func (h *CommunityHandler) CreateCommunity(c *gin.Context) {
	// Garantir que a requisição está usando UTF-8
	c.Header("Content-Type", "application/json; charset=utf-8")

	var community domain.Community
	if err := c.ShouldBindJSON(&community); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	err := h.CommunityService.CreateCommunity(&community)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comunidade criada com sucesso!"})
}

// 🔹 Buscar todos os Comunidadeos
// @Summary Buscar todos os Comunidadeos
// @Description Buscar todos os Comunidadeos
// @Tags Comunidade
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Community
// @Router /community [get]
func (h *CommunityHandler) BuscarComunidades(c *gin.Context) {
	communities, err := h.CommunityService.GetAllCommunits()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar Comunidadeos"})
		return
	}

	c.JSON(http.StatusOK, communities)
}
