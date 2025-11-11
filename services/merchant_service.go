package services

import (
	"errors"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/auth"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/merchant"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/lib/utils"
)

type merchantService struct {
	repo     merchant.Repository
	authRepo auth.Repository
}

// NewMerchantService creates a new merchant service
func NewMerchantService(repo merchant.Repository, authRepo auth.Repository) merchant.Service {
	return &merchantService{
		repo:     repo,
		authRepo: authRepo,
	}
}

// RegisterMerchant registers a new merchant
func (s *merchantService) RegisterMerchant(req *merchant.RegisterMerchantRequest) (*merchant.Merchant, error) {
	// Check if email already exists
	existingUser, _ := s.authRepo.FindUserByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user account
	user := &auth.User{
		Email:    req.Email,
		Password: hashedPassword,
		Name:     req.Name,
		Phone:    req.Phone,
		Role:     "merchant",
		IsActive: true,
	}

	if err := s.authRepo.CreateUser(user); err != nil {
		return nil, err
	}

	// Create merchant profile
	merch := &merchant.Merchant{
		UserID:      user.ID,
		ShopName:    req.ShopName,
		ShopAddress: req.ShopAddress,
		Phone:       req.Phone,
		Latitude:    req.Latitude,
		Longitude:   req.Longitude,
		Description: req.Description,
		IsVerified:  false,
		IsActive:    true,
	}

	if err := s.repo.Create(merch); err != nil {
		return nil, err
	}

	return merch, nil
}

// GetMerchantByID gets a merchant by ID
func (s *merchantService) GetMerchantByID(id uint) (*merchant.Merchant, error) {
	return s.repo.FindByID(id)
}

// GetMerchantByUserID gets a merchant by user ID
func (s *merchantService) GetMerchantByUserID(userID uint) (*merchant.Merchant, error) {
	return s.repo.FindByUserID(userID)
}

// UpdateMerchant updates merchant information
func (s *merchantService) UpdateMerchant(id uint, req *merchant.UpdateMerchantRequest) error {
	// Get existing merchant
	merch, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Update fields
	if req.ShopName != "" {
		merch.ShopName = req.ShopName
	}
	if req.ShopAddress != "" {
		merch.ShopAddress = req.ShopAddress
	}
	if req.Phone != "" {
		merch.Phone = req.Phone
	}
	if req.Latitude != 0 {
		merch.Latitude = req.Latitude
	}
	if req.Longitude != 0 {
		merch.Longitude = req.Longitude
	}
	if req.Description != "" {
		merch.Description = req.Description
	}

	return s.repo.Update(merch)
}
