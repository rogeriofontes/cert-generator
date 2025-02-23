package domain

import "gorm.io/gorm"

// Participante representa um participante em um evento
// @Description Estrutura que define um participante
type Participant struct {
	gorm.Model
	Name          string `json:"name" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
	EventID       uint   `json:"event_id" binding:"required"`
	Event         *Event `json:"event,omitempty" gorm:"foreignKey:EventID"`
	Status        string `json:"status" gorm:"default:'pendente'"`
	Certificate   string `json:"certificate,omitempty"` // Permite que não seja enviado no JSON
	CertificateId string `json:"certificate_id,omitempty"`
}
