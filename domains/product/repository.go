package product

// Repository defines the interface for product data operations
type Repository interface {
	Create(product *Product) error
	FindByID(id uint) (*Product, error)
	FindAll(filter *SearchFilter) ([]Product, int64, error)
	Update(product *Product) error
	Delete(id uint) error
	UpdateStock(id uint, quantity int) error
	FindByMerchantID(merchantID uint) ([]Product, error)
}
