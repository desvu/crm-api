package app

import (
	"context"

	"github.com/qilin/crm-api/internal/handler/micro/service/subscriber/sub_game_store"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/web"
	"github.com/qilin/crm-api/internal/app/container/env"
	"github.com/qilin/crm-api/internal/app/container/event"
	"github.com/qilin/crm-api/internal/app/container/handler"
	"github.com/qilin/crm-api/internal/app/container/pkg"
	"github.com/qilin/crm-api/internal/app/container/repository"
	"github.com/qilin/crm-api/internal/app/container/service"
	"go.uber.org/fx"
)

type App struct {
	fxOptions fx.Option
	grpc      micro.Service
	web       web.Service
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
		handler.NewWeb,
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
		//fx.NopLogger,

		fx.Invoke(
			func(w web.Service, grpc micro.Service) (*App, error) {
				app.web = w
				if err := app.web.Init(); err != nil {
					return nil, err
				}

				app.grpc = grpc
				//app.grpc.Init()

				return app, nil
			},
			sub_game_store.New,
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

	if err := app.web.Run(); err != nil {
		return err
	}

	return nil
}
