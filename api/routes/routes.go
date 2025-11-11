package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewUserRoutes),
	fx.Provide(NewAuthRoutes),
	fx.Provide(NewProductRoutes),
	fx.Provide(NewMerchantRoutes),
	fx.Provide(NewOrderRoutes),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	userRoutes UserRoutes,
	authRoutes AuthRoutes,
	productRoutes ProductRoutes,
	merchantRoutes MerchantRoutes,
	orderRoutes OrderRoutes,
) Routes {
	return Routes{
		userRoutes,
		authRoutes,
		productRoutes,
		merchantRoutes,
		orderRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
