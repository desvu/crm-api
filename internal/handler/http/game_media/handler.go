package game_media

import (
	"bytes"
	"io"

	"github.com/qilin/crm-api/pkg/response"

	"github.com/labstack/echo/v4"
	"github.com/qilin/crm-api/internal/domain/enum/game_media"
	"github.com/qilin/crm-api/internal/domain/service"
)

//swagger:parameters reqUpload
type reqUpload struct {
	// in: path
	// example: 11002485-cb51-4b29-8423-cba43f29f143
	GameID string `param:"game_id"`

	// in: path
	// example: 43
	MediaID uint `param:"media_id"`
}

// swagger:route POST /games/{game_id}/media/{media_id} game_media reqUpload
//
// Upload game media
//
// This endpoint returns a list of extended game structures
//
//     Responses:
//       200: Media
func (h Handler) Upload(c echo.Context) error {
	req := new(reqUpload)
	if err := c.Bind(req); err != nil {
		return err
	}

	file, err := c.FormFile("image")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return err
	}

	err = src.Close()
	if err != nil {
		return err
	}

	game, err := h.GameService.GetByID(c.Request().Context(), req.GameID)
	if err != nil {
		return err
	}

	media, err := h.GameMediaService.Upload(c.Request().Context(), &service.UploadGameMediaData{
		Game:  game,
		ID:    req.MediaID,
		Image: buf.Bytes(),
	})

	if err != nil {
		return err
	}

	return response.New(c, h.view(media))
}

//swagger:parameters reqCreate
type reqCreate struct {
	// in: path
	// example: 11002485-cb51-4b29-8423-cba43f29f143
	GameID string `param:"game_id"`

	// in: body
	// enum: [wideSlider vertical horizontal horizontalSmall largeSingle catalog screenshot description]
	Type string `json:"type"`

	// in: body
	// example: png
	Extension string `json:"extension"`
}

// swagger:route POST /games/{game_id}/media game_media reqCreate
//
// Create game media
//
// This endpoint returns a list of extended game structures
//
//     Responses:
//       200: Media
func (h Handler) Create(c echo.Context) error {
	req := new(reqCreate)
	if err := c.Bind(req); err != nil {
		return err
	}

	game, err := h.GameService.GetByID(c.Request().Context(), c.Param("game_id"))
	if err != nil {
		return err
	}

	media, err := h.GameMediaService.Create(c.Request().Context(), &service.CreateGameMediaData{
		Game:      game,
		Type:      game_media.NewTypeByString(req.Type),
		Extension: req.Extension,
	})

	if err != nil {
		return err
	}

	return response.New(c, h.view(media))
}
