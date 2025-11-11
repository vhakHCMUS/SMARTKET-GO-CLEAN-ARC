package location

// Repository defines the interface for location data operations
type Repository interface {
	Create(location *Location) error
	FindByID(id uint) (*Location, error)
	FindByUserID(userID uint) ([]Location, error)
	Update(location *Location) error
	Delete(id uint) error
	SetDefaultLocation(userID uint, locationID uint) error
}
