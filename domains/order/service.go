package order

// Service defines the interface for order business logic
type Service interface {
	// Order operations
	CreateOrder(userID uint, req *CreateOrderRequest) (*Order, error)
	GetOrderByID(id uint) (*Order, error)
	GetOrderByCode(code string) (*Order, error)
	GetUserOrders(userID uint) ([]Order, error)
	GetMerchantOrders(merchantID uint) ([]Order, error)
	RedeemOrder(merchantID uint, orderCode string) error

	// Cart operations
	AddToCart(userID uint, req *AddToCartRequest) error
	GetCart(userID uint) (*Cart, error)
	UpdateCartItem(userID uint, itemID uint, quantity int) error
	RemoveCartItem(userID uint, itemID uint) error
	ClearCart(userID uint) error
}
