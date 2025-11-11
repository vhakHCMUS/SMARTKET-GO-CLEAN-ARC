package product

import (
	"time"
)

// Product represents a product/smart bag in the system
type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	MerchantID  uint      `json:"merchant_id" gorm:"not null"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Category    string    `json:"category" gorm:"not null"` // vegetables, fruits, meat, bakery, etc.
	OrigPrice   float64   `json:"orig_price" gorm:"not null"`
	SalePrice   float64   `json:"sale_price" gorm:"not null"`
	Discount    float64   `json:"discount"` // percentage
	Stock       int       `json:"stock" gorm:"default:0"`
	Images      string    `json:"images"` // comma-separated URLs
	ExpiryDate  time.Time `json:"expiry_date"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// SearchFilter represents search and filter criteria
type SearchFilter struct {
	Keyword    string
	Category   string
	MinPrice   float64
	MaxPrice   float64
	MerchantID uint
	Location   string
	Limit      int
	Offset     int
}

// CreateProductRequest represents request to create a product
type CreateProductRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Category    string    `json:"category" binding:"required"`
	OrigPrice   float64   `json:"orig_price" binding:"required,gt=0"`
	SalePrice   float64   `json:"sale_price" binding:"required,gt=0"`
	Stock       int       `json:"stock" binding:"required,gte=0"`
	Images      string    `json:"images"`
	ExpiryDate  time.Time `json:"expiry_date" binding:"required"`
}

// UpdateProductRequest represents request to update a product
type UpdateProductRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	OrigPrice   float64   `json:"orig_price"`
	SalePrice   float64   `json:"sale_price"`
	Stock       int       `json:"stock"`
	Images      string    `json:"images"`
	ExpiryDate  time.Time `json:"expiry_date"`
	IsActive    bool      `json:"is_active"`
}
