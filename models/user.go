package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`                 // Hashed before storing
	Role     string `json:"role" gorm:"default:user"` // "admin" or "user"
}
