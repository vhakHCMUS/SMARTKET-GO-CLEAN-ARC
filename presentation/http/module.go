package handlers

import (
	"go.uber.org/fx"
)

// Module exports handlers
var Module = fx.Options(
	fx.Provide(NewAuthHandler),
	fx.Provide(NewProductHandler),
	fx.Provide(NewMerchantHandler),
	fx.Provide(NewOrderHandler),
)
