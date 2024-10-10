package rest

import (
	"context"
	"go_bootstrap/internal/helloworld"

	microservice "github.com/SleepNFire/go-microservice"
	msGin "github.com/SleepNFire/go-microservice/gin"
	"go.uber.org/fx"
)

func Init(lc fx.Lifecycle, he *helloworld.HelloEndpoint) (*microservice.MicroService, error) {
	routers := msGin.NewGinService()

	he.RegisterEndpoint(routers.Public.Engine)

	ms := microservice.NewMicroService(routers)

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				ms.Start()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return ms.Stop()
			},
		},
	)

	return ms, nil
}
