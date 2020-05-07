package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/qilin/crm-api/internal/handler/graph"
	"github.com/qilin/crm-api/internal/handler/http/storefront"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Storefronts *storefront.Handler
	Resolver    *graph.Resolver
}

func New(params Params) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/api/graphql/client", echo.WrapHandler(graph.Playground("/api/graphql")))
	e.POST("/api/graphql", echo.WrapHandler(graph.NewHandler(params.Resolver)))

	api := e.Group("/api/v1")
	// manage storefront templates
	api.GET("/storefronts", params.Storefronts.List)
	api.POST("/storefronts", params.Storefronts.Create)
	api.GET("/storefronts/:id", params.Storefronts.Get)
	api.PUT("/storefronts/:id", params.Storefronts.Update)
	api.POST("/storefronts/:id/activate", params.Storefronts.Activate)
	api.DELETE("/storefronts/:id", params.Storefronts.Delete)

	return e
}
