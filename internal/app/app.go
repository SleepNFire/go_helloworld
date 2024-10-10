package app

import (
	"go.uber.org/fx"
	"go_bootstrap/config"
	"go_bootstrap/internal/helloworld"
	"go_bootstrap/internal/rest"
)

func Init() *fx.App {

	app := fx.New(
		fx.Options(
			fx.Provide(config.Init),
			fx.Provide(helloworld.InitHelloEndpoint),

			fx.Invoke(rest.Init),
		),
	)

	return app
}
