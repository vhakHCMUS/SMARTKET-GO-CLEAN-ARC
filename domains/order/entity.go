package order

import (
	"time"
)

// Order represents a customer order
type Order struct {
	ID              uint        `json:"id" gorm:"primaryKey"`
	UserID          uint        `json:"user_id" gorm:"not null"`
	MerchantID      uint        `json:"merchant_id" gorm:"not null"`
	OrderCode       string      `json:"order_code" gorm:"unique;not null"`
	TotalAmount     float64     `json:"total_amount" gorm:"not null"`
	Status          string      `json:"status" gorm:"default:'pending'"` // pending, confirmed, ready, completed, cancelled
	PaymentMethod   string      `json:"payment_method" gorm:"default:'COD'"`
	PaymentStatus   string      `json:"payment_status" gorm:"default:'unpaid'"` // unpaid, paid
	DeliveryAddress string      `json:"delivery_address"`
	PickupTime      time.Time   `json:"pickup_time"`
	CompletedAt     *time.Time  `json:"completed_at"`
	Notes           string      `json:"notes"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	Items           []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	OrderID     uint    `json:"order_id" gorm:"not null"`
	ProductID   uint    `json:"product_id" gorm:"not null"`
	Quantity    int     `json:"quantity" gorm:"not null"`
	Price       float64 `json:"price" gorm:"not null"`
	Subtotal    float64 `json:"subtotal" gorm:"not null"`
	ProductName string  `json:"product_name"`
}

// Cart represents a shopping cart
type Cart struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	UserID    uint       `json:"user_id" gorm:"not null"`
	Items     []CartItem `json:"items" gorm:"foreignKey:CartID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// CartItem represents an item in the cart
type CartItem struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	CartID     uint      `json:"cart_id" gorm:"not null"`
	ProductID  uint      `json:"product_id" gorm:"not null"`
	MerchantID uint      `json:"merchant_id" gorm:"not null"`
	Quantity   int       `json:"quantity" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// CreateOrderRequest represents request to create an order
type CreateOrderRequest struct {
	MerchantID      uint   `json:"merchant_id" binding:"required"`
	DeliveryAddress string `json:"delivery_address" binding:"required"`
	PaymentMethod   string `json:"payment_method" binding:"required"`
	Notes           string `json:"notes"`
	Items           []struct {
		ProductID uint `json:"product_id" binding:"required"`
		Quantity  int  `json:"quantity" binding:"required,gt=0"`
	} `json:"items" binding:"required,min=1"`
}

// AddToCartRequest represents request to add item to cart
type AddToCartRequest struct {
	ProductID  uint `json:"product_id" binding:"required"`
	MerchantID uint `json:"merchant_id" binding:"required"`
	Quantity   int  `json:"quantity" binding:"required,gt=0"`
}

// UpdateCartItemRequest represents request to update cart item quantity
type UpdateCartItemRequest struct {
	Quantity int `json:"quantity" binding:"required,gte=0"`
}

// RedeemOrderRequest represents request to redeem/complete an order
type RedeemOrderRequest struct {
	OrderCode string `json:"order_code" binding:"required"`
}
