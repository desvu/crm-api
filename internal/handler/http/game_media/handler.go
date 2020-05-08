package game_media

import (
	"bytes"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/qilin/crm-api/internal/domain/enum/game_media"
	"github.com/qilin/crm-api/internal/domain/service"
)

func (h Handler) Upload(c echo.Context) error {
	gameMediaID, err := strconv.ParseUint(c.Param("game_media_id"), 10, 32)
	if err != nil {
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

	game, err := h.GameService.GetByID(c.Request().Context(), c.Param("game_id"))
	if err != nil {
		return err
	}

	media, err := h.GameMediaService.Upload(c.Request().Context(), &service.UploadGameMediaData{
		Game:  game,
		ID:    uint(gameMediaID),
		Image: buf.Bytes(),
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, h.view(media))
}

type reqCreate struct {
	Type      string `json:"type"`
	Extension string `json:"extension"`
}

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

	return c.JSON(http.StatusOK, h.view(media))
}
