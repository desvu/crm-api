package game

import (
	"github.com/qilin/crm-api/pkg/response"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetByFilter(c echo.Context) error {
	data, err := convertGetByFilterRequest(c)
	if err != nil {
		return err
	}

	games, err := h.GameService.GetExByFilter(c.Request().Context(), data)
	if err != nil {
		return err
	}

	return response.New(c, h.viewArray(games))
}

func (h Handler) GetByID(c echo.Context) error {
	game, err := h.GameService.GetExByID(c.Request().Context(), c.Param("game_id"))
	if err != nil {
		return err
	}

	return response.New(c, h.view(game))
}

func (h Handler) Upsert(c echo.Context) error {
	data, err := convertUpsertRequest(c)
	if err != nil {
		return err
	}

	game, err := h.GameService.Upsert(c.Request().Context(), data)
	if err != nil {
		return err
	}

	return response.New(c, h.view(game))
}

func (h Handler) Publish(c echo.Context) error {
	err := h.GameService.Publish(c.Request().Context(), c.Param("game_id"))
	if err != nil {
		return err
	}

	game, err := h.GameService.GetExLastPublishedByID(c.Request().Context(), c.Param("game_id"))
	if err != nil {
		return err
	}

	return response.New(c, h.view(game))
}
