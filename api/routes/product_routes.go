package routes

import (
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/api/middlewares"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/lib"
	handlers "github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/presentation/http"
)

// ProductRoutes struct
type ProductRoutes struct {
	handler        *handlers.ProductHandler
	requestHandler lib.RequestHandler
}

// Setup product routes
func (r ProductRoutes) Setup() {
	api := r.requestHandler.Gin.Group("/api")
	{
		// Public routes
		api.GET("/products/search", r.handler.SearchProducts)
		api.GET("/products/:id", r.handler.GetProduct)

		// Merchant routes (requires authentication + merchant role)
		merchant := api.Group("/merchant")
		merchant.Use(middlewares.AuthMiddleware())
		merchant.Use(middlewares.MerchantMiddleware())
		{
			merchant.POST("/products", r.handler.CreateProduct)
			merchant.GET("/products", r.handler.GetMerchantProducts)
			merchant.PUT("/products/:id", r.handler.UpdateProduct)
			merchant.DELETE("/products/:id", r.handler.DeleteProduct)
		}
	}
}

// NewProductRoutes creates new product routes
func NewProductRoutes(
	handler *handlers.ProductHandler,
	requestHandler lib.RequestHandler,
) ProductRoutes {
	return ProductRoutes{
		handler:        handler,
		requestHandler: requestHandler,
	}
}
