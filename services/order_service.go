package services

import (
	"errors"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/order"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/product"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/lib/utils"
	"time"
)

type orderService struct {
	repo        order.Repository
	productRepo product.Repository
}

// NewOrderService creates a new order service
func NewOrderService(repo order.Repository, productRepo product.Repository) order.Service {
	return &orderService{
		repo:        repo,
		productRepo: productRepo,
	}
}

// CreateOrder creates a new order
func (s *orderService) CreateOrder(userID uint, req *order.CreateOrderRequest) (*order.Order, error) {
	var totalAmount float64
	var orderItems []order.OrderItem

	// Calculate total and prepare order items
	for _, item := range req.Items {
		prod, err := s.productRepo.FindByID(item.ProductID)
		if err != nil {
			return nil, errors.New("product not found")
		}

		// Check stock
		if prod.Stock < item.Quantity {
			return nil, errors.New("insufficient stock for product: " + prod.Name)
		}

		// Check if product belongs to the specified merchant
		if prod.MerchantID != req.MerchantID {
			return nil, errors.New("all products must be from the same merchant")
		}

		subtotal := prod.SalePrice * float64(item.Quantity)
		totalAmount += subtotal

		orderItems = append(orderItems, order.OrderItem{
			ProductID:   item.ProductID,
			Quantity:    item.Quantity,
			Price:       prod.SalePrice,
			Subtotal:    subtotal,
			ProductName: prod.Name,
		})

		// Update stock
		prod.Stock -= item.Quantity
		if err := s.productRepo.Update(prod); err != nil {
			return nil, err
		}
	}

	// Create order
	ord := &order.Order{
		UserID:          userID,
		MerchantID:      req.MerchantID,
		OrderCode:       utils.GenerateOrderCode(),
		TotalAmount:     totalAmount,
		Status:          "pending",
		PaymentMethod:   req.PaymentMethod,
		PaymentStatus:   "unpaid",
		DeliveryAddress: req.DeliveryAddress,
		PickupTime:      time.Now().Add(2 * time.Hour), // Default 2 hours from now
		Notes:           req.Notes,
		Items:           orderItems,
	}

	if err := s.repo.CreateOrder(ord); err != nil {
		return nil, err
	}

	return ord, nil
}

// GetOrderByID gets an order by ID
func (s *orderService) GetOrderByID(id uint) (*order.Order, error) {
	return s.repo.FindOrderByID(id)
}

// GetOrderByCode gets an order by order code
func (s *orderService) GetOrderByCode(code string) (*order.Order, error) {
	return s.repo.FindOrderByCode(code)
}

// GetUserOrders gets all orders for a user
func (s *orderService) GetUserOrders(userID uint) ([]order.Order, error) {
	return s.repo.FindOrdersByUserID(userID)
}

// GetMerchantOrders gets all orders for a merchant
func (s *orderService) GetMerchantOrders(merchantID uint) ([]order.Order, error) {
	return s.repo.FindOrdersByMerchantID(merchantID)
}

// RedeemOrder redeems an order (merchant confirms pickup)
func (s *orderService) RedeemOrder(merchantID uint, orderCode string) error {
	ord, err := s.repo.FindOrderByCode(orderCode)
	if err != nil {
		return err
	}

	// Check if order belongs to the merchant
	if ord.MerchantID != merchantID {
		return errors.New("unauthorized")
	}

	// Check if order is in correct status
	if ord.Status != "pending" && ord.Status != "confirmed" && ord.Status != "ready" {
		return errors.New("order cannot be redeemed in current status")
	}

	// Check pickup time validity (example: within 24 hours)
	if time.Now().After(ord.PickupTime.Add(24 * time.Hour)) {
		return errors.New("order pickup time has expired")
	}

	// Update order status
	now := time.Now()
	ord.Status = "completed"
	ord.PaymentStatus = "paid"
	ord.CompletedAt = &now

	return s.repo.UpdateOrder(ord)
}

// AddToCart adds an item to cart
func (s *orderService) AddToCart(userID uint, req *order.AddToCartRequest) error {
	// Get or create cart
	cart, err := s.repo.FindCartByUserID(userID)
	if err != nil {
		return err
	}

	// Verify product exists
	prod, err := s.productRepo.FindByID(req.ProductID)
	if err != nil {
		return errors.New("product not found")
	}

	// Check stock
	if prod.Stock < req.Quantity {
		return errors.New("insufficient stock")
	}

	// Add item to cart
	cartItem := &order.CartItem{
		CartID:     cart.ID,
		ProductID:  req.ProductID,
		MerchantID: req.MerchantID,
		Quantity:   req.Quantity,
	}

	return s.repo.AddCartItem(cartItem)
}

// GetCart gets user's cart
func (s *orderService) GetCart(userID uint) (*order.Cart, error) {
	return s.repo.FindCartByUserID(userID)
}

// UpdateCartItem updates cart item quantity
func (s *orderService) UpdateCartItem(userID uint, itemID uint, quantity int) error {
	// Get cart
	cart, err := s.repo.FindCartByUserID(userID)
	if err != nil {
		return err
	}

	// Get cart item
	item, err := s.repo.FindCartItemByID(itemID)
	if err != nil {
		return err
	}

	// Verify item belongs to user's cart
	if item.CartID != cart.ID {
		return errors.New("unauthorized")
	}

	// If quantity is 0, remove item
	if quantity == 0 {
		return s.repo.RemoveCartItem(itemID)
	}

	// Verify stock
	prod, err := s.productRepo.FindByID(item.ProductID)
	if err != nil {
		return err
	}

	if prod.Stock < quantity {
		return errors.New("insufficient stock")
	}

	// Update quantity
	item.Quantity = quantity
	return s.repo.UpdateCartItem(item)
}

// RemoveCartItem removes an item from cart
func (s *orderService) RemoveCartItem(userID uint, itemID uint) error {
	// Get cart
	cart, err := s.repo.FindCartByUserID(userID)
	if err != nil {
		return err
	}

	// Get cart item
	item, err := s.repo.FindCartItemByID(itemID)
	if err != nil {
		return err
	}

	// Verify item belongs to user's cart
	if item.CartID != cart.ID {
		return errors.New("unauthorized")
	}

	return s.repo.RemoveCartItem(itemID)
}

// ClearCart clears user's cart
func (s *orderService) ClearCart(userID uint) error {
	cart, err := s.repo.FindCartByUserID(userID)
	if err != nil {
		return err
	}

	return s.repo.ClearCart(cart.ID)
}
