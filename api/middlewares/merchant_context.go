package middlewares

import (
	"net/http"

	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/merchant"

	"github.com/gin-gonic/gin"
)

// MerchantContextMiddleware fetches merchant ID and adds it to context
type MerchantContextMiddleware struct {
	merchantService merchant.Service
}

// NewMerchantContextMiddleware creates a new merchant context middleware
func NewMerchantContextMiddleware(merchantService merchant.Service) *MerchantContextMiddleware {
	return &MerchantContextMiddleware{
		merchantService: merchantService,
	}
}

// Handle adds merchant ID to context if user is a merchant
func (m *MerchantContextMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "merchant" {
			c.JSON(http.StatusForbidden, gin.H{"error": "merchant access required"})
			c.Abort()
			return
		}

		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
			c.Abort()
			return
		}

		// Get merchant profile
		merch, err := m.merchantService.GetMerchantByUserID(userID.(uint))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "merchant profile not found"})
			c.Abort()
			return
		}

		c.Set("merchantID", merch.ID)
		c.Next()
	}
}
