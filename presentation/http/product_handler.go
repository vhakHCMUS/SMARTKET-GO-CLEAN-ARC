package handlers

import (
	"net/http"
	"strconv"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/product"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productService product.Service
}

// NewProductHandler creates a new product handler
func NewProductHandler(productService product.Service) *ProductHandler {
	return &ProductHandler{productService: productService}
}

// CreateProduct creates a new product (merchant only)
// @Summary Create a product
// @Tags products
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body product.CreateProductRequest true "Product details"
// @Success 201 {object} product.Product
// @Router /api/merchant/products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "merchant not authenticated"})
		return
	}

	var req product.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	prod, err := h.productService.CreateProduct(merchantID.(uint), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": prod})
}

// GetProduct gets a product by ID
// @Summary Get product details
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} product.Product
// @Router /api/products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	prod, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": prod})
}

// SearchProducts searches for products
// @Summary Search products
// @Tags products
// @Produce json
// @Param keyword query string false "Search keyword"
// @Param category query string false "Product category"
// @Param min_price query number false "Minimum price"
// @Param max_price query number false "Maximum price"
// @Param merchant_id query int false "Merchant ID"
// @Param limit query int false "Limit" default(20)
// @Param offset query int false "Offset" default(0)
// @Success 200 {array} product.Product
// @Router /api/products/search [get]
func (h *ProductHandler) SearchProducts(c *gin.Context) {
	filter := &product.SearchFilter{
		Keyword:  c.Query("keyword"),
		Category: c.Query("category"),
		Location: c.Query("location"),
	}

	if minPrice := c.Query("min_price"); minPrice != "" {
		if price, err := strconv.ParseFloat(minPrice, 64); err == nil {
			filter.MinPrice = price
		}
	}

	if maxPrice := c.Query("max_price"); maxPrice != "" {
		if price, err := strconv.ParseFloat(maxPrice, 64); err == nil {
			filter.MaxPrice = price
		}
	}

	if merchantID := c.Query("merchant_id"); merchantID != "" {
		if id, err := strconv.ParseUint(merchantID, 10, 32); err == nil {
			filter.MerchantID = uint(id)
		}
	}

	if limit := c.Query("limit"); limit != "" {
		if l, err := strconv.Atoi(limit); err == nil {
			filter.Limit = l
		}
	} else {
		filter.Limit = 20
	}

	if offset := c.Query("offset"); offset != "" {
		if o, err := strconv.Atoi(offset); err == nil {
			filter.Offset = o
		}
	}

	products, total, err := h.productService.SearchProducts(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  products,
		"total": total,
	})
}

// UpdateProduct updates a product (merchant only)
// @Summary Update a product
// @Tags products
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param request body product.UpdateProductRequest true "Product updates"
// @Success 200
// @Router /api/merchant/products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "merchant not authenticated"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	var req product.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.productService.UpdateProduct(uint(id), merchantID.(uint), &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product updated successfully"})
}

// DeleteProduct deletes a product (merchant only)
// @Summary Delete a product
// @Tags products
// @Security BearerAuth
// @Param id path int true "Product ID"
// @Success 200
// @Router /api/merchant/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "merchant not authenticated"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	if err := h.productService.DeleteProduct(uint(id), merchantID.(uint)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted successfully"})
}

// GetMerchantProducts gets all products for a merchant
// @Summary Get merchant's products
// @Tags products
// @Security BearerAuth
// @Produce json
// @Success 200 {array} product.Product
// @Router /api/merchant/products [get]
func (h *ProductHandler) GetMerchantProducts(c *gin.Context) {
	merchantID, exists := c.Get("merchantID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "merchant not authenticated"})
		return
	}

	products, err := h.productService.GetMerchantProducts(merchantID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}
