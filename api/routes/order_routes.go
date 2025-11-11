package routes

import (
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/api/middlewares"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/lib"
	handlers "github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/presentation/http"
)

// OrderRoutes struct
type OrderRoutes struct {
	handler                   *handlers.OrderHandler
	requestHandler            lib.RequestHandler
	merchantContextMiddleware *middlewares.MerchantContextMiddleware
}

// Setup order routes
func (r OrderRoutes) Setup() {
	api := r.requestHandler.Gin.Group("/api")
	api.Use(middlewares.AuthMiddleware())
	{
		// Customer routes
		api.POST("/orders", r.handler.CreateOrder)
		api.GET("/orders", r.handler.GetUserOrders)
		api.GET("/orders/:id", r.handler.GetOrder)

		// Cart routes
		api.GET("/cart", r.handler.GetCart)
		api.POST("/cart/add", r.handler.AddToCart)
		api.PUT("/cart/items/:id", r.handler.UpdateCartItem)
		api.DELETE("/cart/items/:id", r.handler.RemoveCartItem)
		api.POST("/cart/clear", r.handler.ClearCart)

		// Merchant routes
		merchant := api.Group("/merchant")
		merchant.Use(r.merchantContextMiddleware.Handle())
		{
			merchant.GET("/orders", r.handler.GetMerchantOrders)
			merchant.POST("/orders/redeem", r.handler.RedeemOrder)
		}
	}
}

// NewOrderRoutes creates new order routes
func NewOrderRoutes(
	handler *handlers.OrderHandler,
	requestHandler lib.RequestHandler,
	merchantContextMiddleware *middlewares.MerchantContextMiddleware,
) OrderRoutes {
	return OrderRoutes{
		handler:                   handler,
		requestHandler:            requestHandler,
		merchantContextMiddleware: merchantContextMiddleware,
	}
}
