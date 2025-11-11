package location

import (
	"time"
)

// Location represents a user's saved location
type Location struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Address   string    `json:"address" gorm:"not null"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	IsDefault bool      `json:"is_default" gorm:"default:false"`
	Label     string    `json:"label"` // home, work, etc.
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AddLocationRequest represents request to add a location
type AddLocationRequest struct {
	Address   string  `json:"address" binding:"required"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Label     string  `json:"label"`
	IsDefault bool    `json:"is_default"`
}
