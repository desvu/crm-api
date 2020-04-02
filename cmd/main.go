package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/qilin/crm-api/internal/app/repository"
	"github.com/qilin/crm-api/internal/app/service"
	"github.com/qilin/crm-api/internal/config"
	"github.com/qilin/crm-api/internal/env"
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

	repos := repository.New(e.Store)
	_ = service.New(repos)

	//e.POST("/users", userHandler.Create)
	//
	//if err = e.Start(":1323"); err != nil {
	//	e.Logger.Fatal(err)
	//}
}
