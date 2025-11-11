package merchant

// Repository defines the interface for merchant data operations
type Repository interface {
	Create(merchant *Merchant) error
	FindByID(id uint) (*Merchant, error)
	FindByUserID(userID uint) (*Merchant, error)
	Update(merchant *Merchant) error
	Delete(id uint) error
	FindAll() ([]Merchant, error)
}
