package services

import (
	"errors"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/product"
	"time"
)

type productService struct {
	repo product.Repository
}

// NewProductService creates a new product service
func NewProductService(repo product.Repository) product.Service {
	return &productService{repo: repo}
}

// CreateProduct creates a new product
func (s *productService) CreateProduct(merchantID uint, req *product.CreateProductRequest) (*product.Product, error) {
	// Calculate discount
	discount := ((req.OrigPrice - req.SalePrice) / req.OrigPrice) * 100

	prod := &product.Product{
		MerchantID:  merchantID,
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		OrigPrice:   req.OrigPrice,
		SalePrice:   req.SalePrice,
		Discount:    discount,
		Stock:       req.Stock,
		Images:      req.Images,
		ExpiryDate:  req.ExpiryDate,
		IsActive:    true,
	}

	if err := s.repo.Create(prod); err != nil {
		return nil, err
	}

	return prod, nil
}

// GetProductByID gets a product by ID
func (s *productService) GetProductByID(id uint) (*product.Product, error) {
	return s.repo.FindByID(id)
}

// SearchProducts searches for products with filters
func (s *productService) SearchProducts(filter *product.SearchFilter) ([]product.Product, int64, error) {
	// Set default limit if not provided
	if filter.Limit == 0 {
		filter.Limit = 20
	}

	return s.repo.FindAll(filter)
}

// UpdateProduct updates a product
func (s *productService) UpdateProduct(id uint, merchantID uint, req *product.UpdateProductRequest) error {
	// Get existing product
	prod, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Check ownership
	if prod.MerchantID != merchantID {
		return errors.New("unauthorized")
	}

	// Update fields
	if req.Name != "" {
		prod.Name = req.Name
	}
	if req.Description != "" {
		prod.Description = req.Description
	}
	if req.Category != "" {
		prod.Category = req.Category
	}
	if req.OrigPrice > 0 {
		prod.OrigPrice = req.OrigPrice
	}
	if req.SalePrice > 0 {
		prod.SalePrice = req.SalePrice
	}
	if req.Stock >= 0 {
		prod.Stock = req.Stock
	}
	if req.Images != "" {
		prod.Images = req.Images
	}
	if !req.ExpiryDate.IsZero() {
		prod.ExpiryDate = req.ExpiryDate
	}
	prod.IsActive = req.IsActive

	// Recalculate discount
	if prod.OrigPrice > 0 && prod.SalePrice > 0 {
		prod.Discount = ((prod.OrigPrice - prod.SalePrice) / prod.OrigPrice) * 100
	}

	prod.UpdatedAt = time.Now()

	return s.repo.Update(prod)
}

// DeleteProduct deletes a product
func (s *productService) DeleteProduct(id uint, merchantID uint) error {
	// Get existing product
	prod, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Check ownership
	if prod.MerchantID != merchantID {
		return errors.New("unauthorized")
	}

	return s.repo.Delete(id)
}

// GetMerchantProducts gets all products for a merchant
func (s *productService) GetMerchantProducts(merchantID uint) ([]product.Product, error) {
	return s.repo.FindByMerchantID(merchantID)
}

// UpdateStock updates product stock
func (s *productService) UpdateStock(id uint, quantity int) error {
	return s.repo.UpdateStock(id, quantity)
}
