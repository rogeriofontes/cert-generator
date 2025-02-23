package domain

import "gorm.io/gorm"

// Event represents an event
// @Description Estrutura que define um evento
type Event struct {
	gorm.Model
	Name         string        `json:"name"`
	Description  string        `json:"description,omitempty"`
	Date         string        `json:"date,omitempty"`
	Local        string        `json:"local"`
	TotalHours   int           `json:"total_hours"`
	CommunityID  uint          `json:"community_id" binding:"required"`
	Community    *Community    `json:"community,omitempty" gorm:"foreignKey:CommunityID"`
	Participants []Participant `gorm:"foreignKey:EventID"`
}
