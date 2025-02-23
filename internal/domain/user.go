package domain

import "gorm.io/gorm"

// User representa um usuário do sistema
type User struct {
	gorm.Model
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" gorm:"unique;not null" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" gorm:"default:'user'"` // "admin" ou "user"
}
