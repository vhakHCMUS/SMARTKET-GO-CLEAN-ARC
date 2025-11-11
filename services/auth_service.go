package services

import (
	"errors"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/auth"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/lib/utils"
	"time"
)

type authService struct {
	repo auth.Repository
}

// NewAuthService creates a new auth service
func NewAuthService(repo auth.Repository) auth.Service {
	return &authService{repo: repo}
}

// Register registers a new user
func (s *authService) Register(req *auth.RegisterRequest) (*auth.User, error) {
	// Check if user already exists
	existingUser, _ := s.repo.FindUserByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := &auth.User{
		Email:    req.Email,
		Password: hashedPassword,
		Name:     req.Name,
		Phone:    req.Phone,
		Role:     "customer",
		IsActive: true,
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login authenticates a user and returns a login response
func (s *authService) Login(req *auth.LoginRequest) (*auth.LoginResponse, error) {
	// Find user by email
	user, err := s.repo.FindUserByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("account is inactive")
	}

	// Verify password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	// Generate tokens
	accessToken, err := utils.GenerateToken(user.ID, user.Email, user.Role, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateToken(user.ID, user.Email, user.Role, 7*24*time.Hour)
	if err != nil {
		return nil, err
	}

	// Create session
	session := &auth.Session{
		UserID:       user.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(24 * time.Hour),
	}

	if err := s.repo.CreateSession(session); err != nil {
		return nil, err
	}

	return &auth.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         *user,
	}, nil
}

// Logout logs out a user
func (s *authService) Logout(token string) error {
	return s.repo.DeleteSession(token)
}

// GetUserByID gets a user by ID
func (s *authService) GetUserByID(id uint) (*auth.User, error) {
	return s.repo.FindUserByID(id)
}

// UpdateProfile updates user profile
func (s *authService) UpdateProfile(user *auth.User) error {
	return s.repo.UpdateUser(user)
}

// ValidateToken validates a token and returns the user
func (s *authService) ValidateToken(token string) (*auth.User, error) {
	session, err := s.repo.FindSessionByToken(token)
	if err != nil {
		return nil, err
	}

	// Check if token is expired
	if session.ExpiresAt.Before(time.Now()) {
		s.repo.DeleteSession(token)
		return nil, errors.New("token expired")
	}

	return &session.User, nil
}
