package auth

// Repository defines the interface for authentication data operations
type Repository interface {
	// User operations
	CreateUser(user *User) error
	FindUserByEmail(email string) (*User, error)
	FindUserByID(id uint) (*User, error)
	UpdateUser(user *User) error

	// Session operations
	CreateSession(session *Session) error
	FindSessionByToken(token string) (*Session, error)
	DeleteSession(token string) error
	DeleteUserSessions(userID uint) error
}
