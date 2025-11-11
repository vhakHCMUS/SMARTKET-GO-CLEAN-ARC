package auth

// Service defines the interface for authentication business logic
type Service interface {
	Register(req *RegisterRequest) (*User, error)
	Login(req *LoginRequest) (*LoginResponse, error)
	Logout(token string) error
	GetUserByID(id uint) (*User, error)
	UpdateProfile(user *User) error
	ValidateToken(token string) (*User, error)
}
