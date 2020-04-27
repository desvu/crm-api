package app

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/micro/go-micro/v2"
	"github.com/qilin/crm-api/internal/app/container/env"
	"github.com/qilin/crm-api/internal/app/container/event"
	"github.com/qilin/crm-api/internal/app/container/handler"
	"github.com/qilin/crm-api/internal/app/container/pkg"
	"github.com/qilin/crm-api/internal/app/container/repository"
	"github.com/qilin/crm-api/internal/app/container/service"
	"github.com/qilin/crm-api/internal/auth"
	envx "github.com/qilin/crm-api/internal/env"
	"go.uber.org/fx"
)

type App struct {
	fxOptions fx.Option
	grpc      micro.Service
	http      *echo.Echo
}

func New() (*App, error) {
	var app = new(App)

	app.FxProvides(
		env.New,
		repository.New,
		service.New,
		event.New,
		pkg.New,
		handler.New,
		handler.NewGRPC,
		handler.NewHTTP,
	)

	return app, nil
}

func (app *App) FxProvides(ff ...func() fx.Option) {
	options := make([]fx.Option, len(ff))
	for i, f := range ff {
		options[i] = f()
	}
	app.fxOptions = fx.Options(options...)
}

func (app *App) Init() error {
	app.fxOptions = fx.Options(
		app.fxOptions,
		fx.NopLogger,

		fx.Provide(
			func(env *envx.Env) (*auth.Auth, error) {
				return auth.New(&env.Auth)
			},
		),

		fx.Invoke(
			func(http *echo.Echo, grpc micro.Service) (*App, error) {
				app.http = http
				app.grpc = grpc
				return app, nil
			},
		),
	)

	err := fx.New(app.fxOptions).Start(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (app *App) Run() error {
	if err := app.grpc.Server().Init(); err != nil {
		return err
	}

	if err := app.grpc.Server().Start(); err != nil {
		return err
	}

	if err := app.http.Start(":8080"); err != nil {
		return err
	}

	return nil
}
