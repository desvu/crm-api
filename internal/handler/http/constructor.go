package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/qilin/crm-api/internal/auth"
	"github.com/qilin/crm-api/internal/handler/graph"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Resolver *graph.Resolver
	Auth     *auth.Auth
}

func New(params Params) (*echo.Echo, error) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(params.Auth.Middleware)

	// Routes
	params.Auth.InitRoutes(e.Group("/api/v1/auth"))

	e.GET("/api/graphql/client", echo.WrapHandler(graph.Playground("/api/graphql")))
	e.POST("/api/graphql", echo.WrapHandler(graph.NewHandler(params.Resolver)))

	return e, nil
}
