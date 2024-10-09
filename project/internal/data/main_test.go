package data_test

import (
	"context"
	"go_bootstrap/project/internal/data"
	"os"
	"testing"

	"go.uber.org/fx"
	"go_bootstrap/project/config"
)

var Mysql data.MySqlAccessor

func TestMain(m *testing.M) {
	if os.Getenv("FUNCTIONNEL_TEST") != "1" {
		return
	}

	ctx := context.Background()

	testApp := fx.New(
		fx.Options(
			fx.Provide(config.Init),
			fx.Provide(data.NewMySqlAccessor),
			fx.Invoke(func(mysql *data.MySqlAccessor) {
				Mysql = *mysql
			}),
		),
	)

	if err := testApp.Start(ctx); err != nil {
		os.Exit(1)
	}
	defer testApp.Stop(ctx)

	code := m.Run()

	os.Exit(code)
}
