package merchant

// Service defines the interface for merchant business logic
type Service interface {
	RegisterMerchant(req *RegisterMerchantRequest) (*Merchant, error)
	GetMerchantByID(id uint) (*Merchant, error)
	GetMerchantByUserID(userID uint) (*Merchant, error)
	UpdateMerchant(id uint, req *UpdateMerchantRequest) error
}
