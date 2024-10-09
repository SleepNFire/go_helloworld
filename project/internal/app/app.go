package app

import (
	"go.uber.org/fx"
	"go_bootstrap/project/config"
	"go_bootstrap/project/internal/data"
	"go_bootstrap/project/internal/rest"
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
