package postgres

import (
	"errors"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/product"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new instance of product repository
func NewProductRepository(db *gorm.DB) product.Repository {
	return &productRepository{db: db}
}

// Create creates a new product
func (r *productRepository) Create(prod *product.Product) error {
	return r.db.Create(prod).Error
}

// FindByID finds a product by ID
func (r *productRepository) FindByID(id uint) (*product.Product, error) {
	var prod product.Product
	err := r.db.First(&prod, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return &prod, nil
}

// FindAll finds all products with filters
func (r *productRepository) FindAll(filter *product.SearchFilter) ([]product.Product, int64, error) {
	var products []product.Product
	var total int64

	query := r.db.Model(&product.Product{}).Where("is_active = ?", true)

	// Apply filters
	if filter.Keyword != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")
	}
	if filter.Category != "" {
		query = query.Where("category = ?", filter.Category)
	}
	if filter.MinPrice > 0 {
		query = query.Where("sale_price >= ?", filter.MinPrice)
	}
	if filter.MaxPrice > 0 {
		query = query.Where("sale_price <= ?", filter.MaxPrice)
	}
	if filter.MerchantID > 0 {
		query = query.Where("merchant_id = ?", filter.MerchantID)
	}

	// Count total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}

	// Order by created_at desc
	query = query.Order("created_at DESC")

	err := query.Find(&products).Error
	if err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

// Update updates a product
func (r *productRepository) Update(prod *product.Product) error {
	return r.db.Save(prod).Error
}

// Delete deletes a product (soft delete by setting is_active to false)
func (r *productRepository) Delete(id uint) error {
	return r.db.Model(&product.Product{}).Where("id = ?", id).Update("is_active", false).Error
}

// UpdateStock updates product stock
func (r *productRepository) UpdateStock(id uint, quantity int) error {
	return r.db.Model(&product.Product{}).Where("id = ?", id).Update("stock", quantity).Error
}

// FindByMerchantID finds all products by merchant ID
func (r *productRepository) FindByMerchantID(merchantID uint) ([]product.Product, error) {
	var products []product.Product
	err := r.db.Where("merchant_id = ?", merchantID).Order("created_at DESC").Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}
