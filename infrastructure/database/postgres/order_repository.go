package postgres

import (
	"errors"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/order"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates a new instance of order repository
func NewOrderRepository(db *gorm.DB) order.Repository {
	return &orderRepository{db: db}
}

// CreateOrder creates a new order
func (r *orderRepository) CreateOrder(ord *order.Order) error {
	return r.db.Create(ord).Error
}

// FindOrderByID finds an order by ID
func (r *orderRepository) FindOrderByID(id uint) (*order.Order, error) {
	var ord order.Order
	err := r.db.Preload("Items").First(&ord, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &ord, nil
}

// FindOrderByCode finds an order by order code
func (r *orderRepository) FindOrderByCode(code string) (*order.Order, error) {
	var ord order.Order
	err := r.db.Preload("Items").Where("order_code = ?", code).First(&ord).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &ord, nil
}

// FindOrdersByUserID finds all orders by user ID
func (r *orderRepository) FindOrdersByUserID(userID uint) ([]order.Order, error) {
	var orders []order.Order
	err := r.db.Preload("Items").Where("user_id = ?", userID).Order("created_at DESC").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// FindOrdersByMerchantID finds all orders by merchant ID
func (r *orderRepository) FindOrdersByMerchantID(merchantID uint) ([]order.Order, error) {
	var orders []order.Order
	err := r.db.Preload("Items").Where("merchant_id = ?", merchantID).Order("created_at DESC").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// UpdateOrder updates an order
func (r *orderRepository) UpdateOrder(ord *order.Order) error {
	return r.db.Save(ord).Error
}

// CreateCart creates a new cart
func (r *orderRepository) CreateCart(cart *order.Cart) error {
	return r.db.Create(cart).Error
}

// FindCartByUserID finds a cart by user ID
func (r *orderRepository) FindCartByUserID(userID uint) (*order.Cart, error) {
	var cart order.Cart
	err := r.db.Preload("Items").Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create a new cart if not found
			newCart := &order.Cart{UserID: userID}
			if err := r.CreateCart(newCart); err != nil {
				return nil, err
			}
			return newCart, nil
		}
		return nil, err
	}
	return &cart, nil
}

// AddCartItem adds an item to cart
func (r *orderRepository) AddCartItem(item *order.CartItem) error {
	// Check if item already exists
	var existing order.CartItem
	err := r.db.Where("cart_id = ? AND product_id = ?", item.CartID, item.ProductID).First(&existing).Error
	if err == nil {
		// Update quantity if exists
		existing.Quantity += item.Quantity
		return r.db.Save(&existing).Error
	}
	return r.db.Create(item).Error
}

// UpdateCartItem updates a cart item
func (r *orderRepository) UpdateCartItem(item *order.CartItem) error {
	return r.db.Save(item).Error
}

// RemoveCartItem removes an item from cart
func (r *orderRepository) RemoveCartItem(id uint) error {
	return r.db.Delete(&order.CartItem{}, id).Error
}

// ClearCart clears all items from cart
func (r *orderRepository) ClearCart(cartID uint) error {
	return r.db.Where("cart_id = ?", cartID).Delete(&order.CartItem{}).Error
}

// FindCartItemsByCartID finds all items in a cart
func (r *orderRepository) FindCartItemsByCartID(cartID uint) ([]order.CartItem, error) {
	var items []order.CartItem
	err := r.db.Where("cart_id = ?", cartID).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

// FindCartItemByID finds a cart item by ID
func (r *orderRepository) FindCartItemByID(id uint) (*order.CartItem, error) {
	var item order.CartItem
	err := r.db.First(&item, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("cart item not found")
		}
		return nil, err
	}
	return &item, nil
}
