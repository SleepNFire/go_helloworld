package app

import (
	"github.com/SleepNFire/go_bootstrap/project/config"
	"github.com/SleepNFire/go_bootstrap/project/internal/data"
	"github.com/SleepNFire/go_bootstrap/project/internal/rest"
	"go.uber.org/fx"
)

func Init() *fx.App {

	app := fx.New(
		fx.Options(
			fx.Provide(config.Init),
			fx.Provide(data.NewMySqlAccessor),
			fx.Invoke(rest.Init),
		),
	)

	return app
}
