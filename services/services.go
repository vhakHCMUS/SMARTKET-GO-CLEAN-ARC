package services

import "go.uber.org/fx"

// Module exports services present
var Module = fx.Options(
	fx.Provide(NewUserService),
	fx.Provide(NewJWTAuthService),
	fx.Provide(NewAuthService),
	fx.Provide(NewProductService),
	fx.Provide(NewMerchantService),
	fx.Provide(NewOrderService),
	fx.Provide(NewLocationService),
)
