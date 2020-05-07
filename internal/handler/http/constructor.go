package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/qilin/crm-api/internal/handler/http/game"
	"github.com/qilin/crm-api/internal/handler/http/game_media"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	GameMediaHandler game_media.Handler
	GameHandler      game.Handler
}

func New(params Params) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/api/games", params.GameHandler.Upsert)
	e.GET("/api/games/:game_id", params.GameHandler.GetByID)

	e.POST("/api/games/:game_id/media", params.GameMediaHandler.Create)
	e.PUT("/api/games/:game_id/media/:game_media_id", params.GameMediaHandler.Upload)

	return e
}
