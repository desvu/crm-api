package app

import (
	"context"

	"github.com/micro/go-micro/web"
	"github.com/qilin/crm-api/internal/app/container/env"
	"github.com/qilin/crm-api/internal/app/container/handler"
	"github.com/qilin/crm-api/internal/app/container/repository"
	"github.com/qilin/crm-api/internal/app/container/service"
	"github.com/qilin/crm-api/internal/app/container/transactor"
	"github.com/qilin/crm-api/internal/handler/micro"
	"go.uber.org/fx"
)

type App struct {
	fxOptions fx.Option
	server    web.Service
}

func New() (*App, error) {
	var app = new(App)

	app.FxProvides(
		env.New,
		repository.New,
		service.New,
		transactor.New,
		handler.New,
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

		fx.Invoke(
			func(params micro.Params) (web.Service, error) {
				var err error
				app.server, err = micro.New(params)
				if err != nil {
					return nil, err
				}

				return app.server, nil
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
	return app.server.Run()
}
