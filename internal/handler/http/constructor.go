// Package classification Qilin CRM API
//
// Documentation of Qilin CRM api.
//
//     Schemes: http
//     BasePath: /api/v1
//     Version: 1
//     Host: localhost:7002
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/qilin/crm-api/internal/handler/http/developer"
	"github.com/qilin/crm-api/internal/handler/http/game"
	"github.com/qilin/crm-api/internal/handler/http/game_media"
	"github.com/qilin/crm-api/internal/handler/http/storefront"
	"github.com/qilin/crm-api/pkg/response"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Storefronts      *storefront.Handler
	GameMediaHandler game_media.Handler
	GameHandler      game.Handler
	DeveloperHandler developer.Handler
}

func New(params Params) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.HTTPErrorHandler = response.Err
	e.Validator = NewValidator()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api/v1")

	// developers
	api.GET("/developers", params.DeveloperHandler.List)

	// manage games
	api.POST("/games", params.GameHandler.Upsert)
	api.GET("/games/:game_id", params.GameHandler.GetByID)
	api.GET("/games", params.GameHandler.GetByFilter)
	api.POST("/games/:game_id/publish", params.GameHandler.Publish)

	// media files upload
	api.POST("/games/:game_id/media", params.GameMediaHandler.Create)
	api.PUT("/games/:game_id/media/:media_id", params.GameMediaHandler.Upload)

	// manage storefront templates
	api.GET("/storefronts", params.Storefronts.List)
	api.POST("/storefronts", params.Storefronts.Create)
	api.GET("/storefronts/:id", params.Storefronts.Get)
	api.PUT("/storefronts/:id", params.Storefronts.Update)
	api.POST("/storefronts/:id/activate", params.Storefronts.Activate)
	api.DELETE("/storefronts/:id", params.Storefronts.Delete)

	return e
}
