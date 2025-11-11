package postgres

import (
	"errors"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/location"
	"gorm.io/gorm"
)

type locationRepository struct {
	db *gorm.DB
}

// NewLocationRepository creates a new instance of location repository
func NewLocationRepository(db *gorm.DB) location.Repository {
	return &locationRepository{db: db}
}

// Create creates a new location
func (r *locationRepository) Create(loc *location.Location) error {
	return r.db.Create(loc).Error
}

// FindByID finds a location by ID
func (r *locationRepository) FindByID(id uint) (*location.Location, error) {
	var loc location.Location
	err := r.db.First(&loc, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("location not found")
		}
		return nil, err
	}
	return &loc, nil
}

// FindByUserID finds all locations by user ID
func (r *locationRepository) FindByUserID(userID uint) ([]location.Location, error) {
	var locations []location.Location
	err := r.db.Where("user_id = ?", userID).Order("is_default DESC, created_at DESC").Find(&locations).Error
	if err != nil {
		return nil, err
	}
	return locations, nil
}

// Update updates a location
func (r *locationRepository) Update(loc *location.Location) error {
	return r.db.Save(loc).Error
}

// Delete deletes a location
func (r *locationRepository) Delete(id uint) error {
	return r.db.Delete(&location.Location{}, id).Error
}

// SetDefaultLocation sets a location as default and unsets others
func (r *locationRepository) SetDefaultLocation(userID uint, locationID uint) error {
	// Start a transaction
	return r.db.Transaction(func(tx *gorm.DB) error {
		// Unset all default locations for this user
		if err := tx.Model(&location.Location{}).Where("user_id = ?", userID).Update("is_default", false).Error; err != nil {
			return err
		}
		// Set the specified location as default
		if err := tx.Model(&location.Location{}).Where("id = ? AND user_id = ?", locationID, userID).Update("is_default", true).Error; err != nil {
			return err
		}
		return nil
	})
}
