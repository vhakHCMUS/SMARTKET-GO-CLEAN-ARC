package postgres

import (
	"errors"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/merchant"
	"gorm.io/gorm"
)

type merchantRepository struct {
	db *gorm.DB
}

// NewMerchantRepository creates a new instance of merchant repository
func NewMerchantRepository(db *gorm.DB) merchant.Repository {
	return &merchantRepository{db: db}
}

// Create creates a new merchant
func (r *merchantRepository) Create(merch *merchant.Merchant) error {
	return r.db.Create(merch).Error
}

// FindByID finds a merchant by ID
func (r *merchantRepository) FindByID(id uint) (*merchant.Merchant, error) {
	var merch merchant.Merchant
	err := r.db.First(&merch, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("merchant not found")
		}
		return nil, err
	}
	return &merch, nil
}

// FindByUserID finds a merchant by user ID
func (r *merchantRepository) FindByUserID(userID uint) (*merchant.Merchant, error) {
	var merch merchant.Merchant
	err := r.db.Where("user_id = ?", userID).First(&merch).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("merchant not found")
		}
		return nil, err
	}
	return &merch, nil
}

// Update updates a merchant
func (r *merchantRepository) Update(merch *merchant.Merchant) error {
	return r.db.Save(merch).Error
}

// Delete deletes a merchant (soft delete by setting is_active to false)
func (r *merchantRepository) Delete(id uint) error {
	return r.db.Model(&merchant.Merchant{}).Where("id = ?", id).Update("is_active", false).Error
}

// FindAll finds all active merchants
func (r *merchantRepository) FindAll() ([]merchant.Merchant, error) {
	var merchants []merchant.Merchant
	err := r.db.Where("is_active = ?", true).Find(&merchants).Error
	if err != nil {
		return nil, err
	}
	return merchants, nil
}
