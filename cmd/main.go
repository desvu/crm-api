package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/qilin/crm-api/internal/config"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/internal/graphql"
	gamerepository "github.com/qilin/crm-api/internal/repository/game"
	gameservice "github.com/qilin/crm-api/internal/service/game"
)

func main() {
	srv := echo.New()
	srv.Logger.SetLevel(log.DEBUG)
	ctx := context.Background()

	cfg, err := config.New()
	if err != nil {
		srv.Logger.Fatal(err)
	}

	e, err := env.New(ctx, cfg)
	if err != nil {
		srv.Logger.Fatal(err)
	}

	gameRepository := gamerepository.New(e.Store)
	gameService := gameservice.New(gameservice.ServiceParams{GameRepository: gameRepository})

	// register graphql api handlers
	gqlResolver := graphql.NewResolver(gameService)
	srv.Any("/api/graphql/client", echo.WrapHandler(graphql.Playground("/api/graphql")))
	srv.Any("/api/graphql", echo.WrapHandler(graphql.NewHandler(gqlResolver)))

	if err = srv.Start(":8080"); err != nil {
		srv.Logger.Fatal(err)
	}
}
