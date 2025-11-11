package merchant

import (
	"time"
)

// Merchant represents a shop/store owner
type Merchant struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id" gorm:"unique;not null"`
	ShopName    string    `json:"shop_name" gorm:"not null"`
	ShopAddress string    `json:"shop_address" gorm:"not null"`
	Phone       string    `json:"phone" gorm:"not null"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Description string    `json:"description"`
	IsVerified  bool      `json:"is_verified" gorm:"default:false"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// RegisterMerchantRequest represents merchant registration data
type RegisterMerchantRequest struct {
	Email       string  `json:"email" binding:"required,email"`
	Password    string  `json:"password" binding:"required,min=6"`
	Name        string  `json:"name" binding:"required"`
	ShopName    string  `json:"shop_name" binding:"required"`
	ShopAddress string  `json:"shop_address" binding:"required"`
	Phone       string  `json:"phone" binding:"required"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Description string  `json:"description"`
}

// UpdateMerchantRequest represents merchant update data
type UpdateMerchantRequest struct {
	ShopName    string  `json:"shop_name"`
	ShopAddress string  `json:"shop_address"`
	Phone       string  `json:"phone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Description string  `json:"description"`
}
