package services

import (
	"errors"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/location"
)

type locationService struct {
	repo location.Repository
}

// NewLocationService creates a new location service
func NewLocationService(repo location.Repository) location.Service {
	return &locationService{repo: repo}
}

// AddLocation adds a new location for a user
func (s *locationService) AddLocation(userID uint, req *location.AddLocationRequest) (*location.Location, error) {
	loc := &location.Location{
		UserID:    userID,
		Address:   req.Address,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Label:     req.Label,
		IsDefault: req.IsDefault,
	}

	if err := s.repo.Create(loc); err != nil {
		return nil, err
	}

	// If this location is set as default, update others
	if req.IsDefault {
		s.repo.SetDefaultLocation(userID, loc.ID)
	}

	return loc, nil
}

// GetUserLocations gets all locations for a user
func (s *locationService) GetUserLocations(userID uint) ([]location.Location, error) {
	return s.repo.FindByUserID(userID)
}

// SetDefaultLocation sets a location as default
func (s *locationService) SetDefaultLocation(userID uint, locationID uint) error {
	return s.repo.SetDefaultLocation(userID, locationID)
}

// DeleteLocation deletes a location
func (s *locationService) DeleteLocation(userID uint, locationID uint) error {
	// Get location to verify ownership
	loc, err := s.repo.FindByID(locationID)
	if err != nil {
		return err
	}

	if loc.UserID != userID {
		return errors.New("unauthorized")
	}

	return s.repo.Delete(locationID)
}
