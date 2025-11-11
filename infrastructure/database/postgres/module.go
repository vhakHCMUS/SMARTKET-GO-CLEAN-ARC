package postgres

import (
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/auth"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/location"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/merchant"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/order"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/domains/product"
	"go.uber.org/fx"
)

// Module exports repository implementations
var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			NewAuthRepository,
			fx.As(new(auth.Repository)),
		),
	),
	fx.Provide(
		fx.Annotate(
			NewProductRepository,
			fx.As(new(product.Repository)),
		),
	),
	fx.Provide(
		fx.Annotate(
			NewMerchantRepository,
			fx.As(new(merchant.Repository)),
		),
	),
	fx.Provide(
		fx.Annotate(
			NewOrderRepository,
			fx.As(new(order.Repository)),
		),
	),
	fx.Provide(
		fx.Annotate(
			NewLocationRepository,
			fx.As(new(location.Repository)),
		),
	),
)
