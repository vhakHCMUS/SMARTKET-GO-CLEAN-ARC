package product

// Service defines the interface for product business logic
type Service interface {
	CreateProduct(merchantID uint, req *CreateProductRequest) (*Product, error)
	GetProductByID(id uint) (*Product, error)
	SearchProducts(filter *SearchFilter) ([]Product, int64, error)
	UpdateProduct(id uint, merchantID uint, req *UpdateProductRequest) error
	DeleteProduct(id uint, merchantID uint) error
	GetMerchantProducts(merchantID uint) ([]Product, error)
	UpdateStock(id uint, quantity int) error
}
