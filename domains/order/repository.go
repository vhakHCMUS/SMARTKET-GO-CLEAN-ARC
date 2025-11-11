package order

// Repository defines the interface for order data operations
type Repository interface {
	// Order operations
	CreateOrder(order *Order) error
	FindOrderByID(id uint) (*Order, error)
	FindOrderByCode(code string) (*Order, error)
	FindOrdersByUserID(userID uint) ([]Order, error)
	FindOrdersByMerchantID(merchantID uint) ([]Order, error)
	UpdateOrder(order *Order) error

	// Cart operations
	CreateCart(cart *Cart) error
	FindCartByUserID(userID uint) (*Cart, error)
	AddCartItem(item *CartItem) error
	UpdateCartItem(item *CartItem) error
	RemoveCartItem(id uint) error
	ClearCart(cartID uint) error
	FindCartItemsByCartID(cartID uint) ([]CartItem, error)
	FindCartItemByID(id uint) (*CartItem, error)
}
