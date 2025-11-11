package bootstrap

import (
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/api/controllers"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/api/middlewares"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/api/routes"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/infrastructure/database/postgres"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/lib"
	handlers "github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/presentation/http"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/repository"
	"github.com/vhakHCMUS/SMARTKET-GO-CLEAN-ARC/services"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	middlewares.Module,
	repository.Module,
	postgres.Module,
	handlers.Module,
	fx.Provide(middlewares.NewMerchantContextMiddleware),
)
