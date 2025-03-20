package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Name     string  `json:"name" gorm:"not null"`
	Address  string  `json:"address"`
	Cuisine  string  `json:"cuisine"`
	Rating   float32 `json:"rating"` // limit to 0.0-5.0 (inclusive)?
	ImageURL string  `json:"imageurl"`
}
