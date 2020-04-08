package main

import (
	"context"

	"github.com/qilin/crm-api/pkg/transactor"

	"github.com/labstack/echo/v4"
	"github.com/qilin/crm-api/internal/app/repository"
	"github.com/qilin/crm-api/internal/app/service"
	"github.com/qilin/crm-api/internal/config"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/graphql"
)

func main() {
	srv := echo.New()
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		srv.Logger.Fatal(err)
	}

	tx, txStore := transactor.New()
	e, err := env.New(ctx, cfg, txStore)
	if err != nil {
		srv.Logger.Fatal(err)
	}

	repos := repository.New(e.Store)
	services := service.New(repos, tx)

	// register graphql api handlers
	gqlResolver := graphql.NewResolver(services.GameService)
	srv.Any("/api/graphql/client", echo.WrapHandler(graphql.Playground("/api/graphql")))
	srv.Any("/api/graphql", echo.WrapHandler(graphql.NewHandler(gqlResolver)))

	if err = srv.Start(":8080"); err != nil {
		srv.Logger.Fatal(err)
	}
}
