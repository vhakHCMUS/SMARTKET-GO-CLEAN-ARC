package postgres

import (
	"errors"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/auth"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

// NewAuthRepository creates a new instance of auth repository
func NewAuthRepository(db *gorm.DB) auth.Repository {
	return &authRepository{db: db}
}

// CreateUser creates a new user
func (r *authRepository) CreateUser(user *auth.User) error {
	return r.db.Create(user).Error
}

// FindUserByEmail finds a user by email
func (r *authRepository) FindUserByEmail(email string) (*auth.User, error) {
	var user auth.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// FindUserByID finds a user by ID
func (r *authRepository) FindUserByID(id uint) (*auth.User, error) {
	var user auth.User
	err := r.db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a user
func (r *authRepository) UpdateUser(user *auth.User) error {
	return r.db.Save(user).Error
}

// CreateSession creates a new session
func (r *authRepository) CreateSession(session *auth.Session) error {
	return r.db.Create(session).Error
}

// FindSessionByToken finds a session by access token
func (r *authRepository) FindSessionByToken(token string) (*auth.Session, error) {
	var session auth.Session
	err := r.db.Preload("User").Where("access_token = ?", token).First(&session).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("session not found")
		}
		return nil, err
	}
	return &session, nil
}

// DeleteSession deletes a session by token
func (r *authRepository) DeleteSession(token string) error {
	return r.db.Where("access_token = ?", token).Delete(&auth.Session{}).Error
}

// DeleteUserSessions deletes all sessions for a user
func (r *authRepository) DeleteUserSessions(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&auth.Session{}).Error
}
