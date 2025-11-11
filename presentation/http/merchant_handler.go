package handlers

import (
	"net/http"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/auth"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/merchant"

	"github.com/gin-gonic/gin"
)

type MerchantHandler struct {
	merchantService merchant.Service
	authService     auth.Service
}

// NewMerchantHandler creates a new merchant handler
func NewMerchantHandler(merchantService merchant.Service, authService auth.Service) *MerchantHandler {
	return &MerchantHandler{
		merchantService: merchantService,
		authService:     authService,
	}
}

// RegisterMerchant handles merchant registration
// @Summary Register a new merchant
// @Tags merchant
// @Accept json
// @Produce json
// @Param request body merchant.RegisterMerchantRequest true "Merchant registration details"
// @Success 201 {object} merchant.Merchant
// @Router /api/merchant/register [post]
func (h *MerchantHandler) RegisterMerchant(c *gin.Context) {
	var req merchant.RegisterMerchantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	merch, err := h.merchantService.RegisterMerchant(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": merch})
}

// LoginMerchant handles merchant login
// @Summary Login merchant
// @Tags merchant
// @Accept json
// @Produce json
// @Param request body auth.LoginRequest true "Login credentials"
// @Success 200 {object} auth.LoginResponse
// @Router /api/merchant/login [post]
func (h *MerchantHandler) LoginMerchant(c *gin.Context) {
	var req auth.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.authService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Verify user is a merchant
	if response.User.Role != "merchant" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not a merchant account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

// GetMerchantProfile gets merchant profile
// @Summary Get merchant profile
// @Tags merchant
// @Security BearerAuth
// @Produce json
// @Success 200 {object} merchant.Merchant
// @Router /api/merchant/profile [get]
func (h *MerchantHandler) GetMerchantProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "merchant not authenticated"})
		return
	}

	merch, err := h.merchantService.GetMerchantByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "merchant not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": merch})
}

// UpdateMerchantProfile updates merchant profile
// @Summary Update merchant profile
// @Tags merchant
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body merchant.UpdateMerchantRequest true "Update details"
// @Success 200
// @Router /api/merchant/profile [put]
func (h *MerchantHandler) UpdateMerchantProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "merchant not authenticated"})
		return
	}

	merch, err := h.merchantService.GetMerchantByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "merchant not found"})
		return
	}

	var req merchant.UpdateMerchantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.merchantService.UpdateMerchant(merch.ID, &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "merchant profile updated successfully"})
}
