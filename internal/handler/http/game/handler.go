package game

import (
	"github.com/qilin/crm-api/pkg/response"

	"github.com/labstack/echo/v4"
)

// swagger:route GET /games games
//
// Getting a list of games by filter
//
// This endpoint returns a list of extended game structures
//
//     Responses:
//       200: GameList
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

// swagger:route GET /games/{id} games
//
// Getting a game by ID
//
// This endpoint returns the extended structure of the game
//
//     Responses:
//       200: Game
func (h Handler) GetByID(c echo.Context) error {
	game, err := h.GameService.GetExByID(c.Request().Context(), c.Param("game_id"))
	if err != nil {
		return err
	}

	return response.New(c, h.view(game))
}

// swagger:route POST /games games reqUpsert
//
// Create or update game information
//
// This endpoint will create or update the game and return the extended game structure
//
//     Responses:
//       200: Game
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

// swagger:route POST /games/{id}/publish games
//
// Publishing the game to the store
//
// This endpoint will publish the game to the store and return the extended game structure
//
//     Responses:
//       200: Game
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
