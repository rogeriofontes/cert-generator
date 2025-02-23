package domain

import "gorm.io/gorm"

// Community representa uma comunidade
// @Description Estrutura que define uma comunidade
type Community struct {
	gorm.Model
	Name      string `json:"name"`
	Organizer string `json:"organizer"`
}
