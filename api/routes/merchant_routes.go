package routes

import (
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/api/middlewares"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/lib"
	handlers "github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/presentation/http"
)

// MerchantRoutes struct
type MerchantRoutes struct {
	handler        *handlers.MerchantHandler
	requestHandler lib.RequestHandler
}

// Setup merchant routes
func (r MerchantRoutes) Setup() {
	api := r.requestHandler.Gin.Group("/api/merchant")
	{
		// Public routes
		api.POST("/register", r.handler.RegisterMerchant)
		api.POST("/login", r.handler.LoginMerchant)

		// Protected routes
		auth := api.Group("")
		auth.Use(middlewares.AuthMiddleware())
		auth.Use(middlewares.MerchantMiddleware())
		{
			auth.GET("/profile", r.handler.GetMerchantProfile)
			auth.PUT("/profile", r.handler.UpdateMerchantProfile)
		}
	}
}

// NewMerchantRoutes creates new merchant routes
func NewMerchantRoutes(
	handler *handlers.MerchantHandler,
	requestHandler lib.RequestHandler,
) MerchantRoutes {
	return MerchantRoutes{
		handler:        handler,
		requestHandler: requestHandler,
	}
}
