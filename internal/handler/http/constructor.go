package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/qilin/crm-api/internal/handler/graph"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Resolver *graph.Resolver
}

func New(params Params) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/api/graphql/client", echo.WrapHandler(graph.Playground("/api/graphql")))
	e.POST("/api/graphql", echo.WrapHandler(graph.NewHandler(params.Resolver)))

	return e
}
