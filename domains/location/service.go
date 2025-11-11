package location

// Service defines the interface for location business logic
type Service interface {
	AddLocation(userID uint, req *AddLocationRequest) (*Location, error)
	GetUserLocations(userID uint) ([]Location, error)
	SetDefaultLocation(userID uint, locationID uint) error
	DeleteLocation(userID uint, locationID uint) error
}
