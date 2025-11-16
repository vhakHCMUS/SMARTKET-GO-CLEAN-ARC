package lib

import (
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewRequestHandler),
	fx.Provide(NewEnv),
	fx.Provide(GetLogger),
	fx.Provide(NewDatabase),
	fx.Provide(func(db Database) *gorm.DB {
		return db.DB
	}),
)
