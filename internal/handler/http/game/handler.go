package game

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetByID(c echo.Context) error {
	game, err := h.GameService.GetExByID(c.Request().Context(), c.Param("game_id"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, h.view(game))
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

	return c.JSON(http.StatusOK, h.view(game))
}
