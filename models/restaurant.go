package models

type Restaurant struct {
	ID       uint    `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name" gorm:"not null"`
	Location string  `json:"location"`
	Cuisine  string  `json:"cuisine"`
	Rating   float32 `json:"rating"` // limit to 0.0-5.0 (inclusive)?
}
