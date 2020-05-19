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
	// required: true
	// example: 11002485-cb51-4b29-8423-cba43f29f143
	GameID string `param:"game_id"`

	// in: path
	// required: true
	// example: 43
	MediaID uint `param:"media_id"`

	// in: formData
	// swagger:file image
	File interface{}
}

// swagger:route PUT /games/{game_id}/media/{media_id} game_media reqUpload
//
// Uploads game media images
//
//	Поддерживается загрузка только png изображение, которые будут перекодированны
//	в jpg изображения. Размер изображения будет изменен до минимального необходимого
// 	разрешения.
//
//	В зависимости от типа изображения, к нему выставляются определенные тробования:
//	wideSlider:
//		необходимое соотношение сторон 16:9
//		минимальной разрешение 1064 * 560
// 	vertical:
//		необходимое соотношение сторон 3:4
//		минимальной разрешение 200 * 266
// 	horizontal:
//		необходимое соотношение сторон 16:9
//		минимальной разрешение 524 * 294
// 	horizontalSmall:
//		необходимое соотношение сторон 16:9
//		минимальной разрешение 254 * 142
// 	largeSingle:
//		необходимое соотношение сторон 16:9
//		минимальной разрешение 744 * 410
// 	catalog:
//		необходимое соотношение сторон 16:9
//		минимальной разрешение 88 * 50
// 	screenshot:
//		необходимое соотношение сторон 16:9
//		минимальной разрешение 1064 * 562
// 	description:
//		любое соотношение сторон
//		любое разрешение
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
	// required: true
	// example: 11002485-cb51-4b29-8423-cba43f29f143
	GameID string `param:"game_id"`

	// in: body
	Body reqCreateBody
}

type reqCreateBody struct {
	// required: true
	// enum: wideSlider,vertical,horizontal,horizontalSmall,largeSingle,catalog,screenshot,description
	Type string `json:"type"`
}

// swagger:route POST /games/{game_id}/media game_media reqCreate
//
// Create
//
// This endpoint returns a list of extended game structures
//
//     Responses:
//       200: Media
func (h Handler) Create(c echo.Context) error {
	//req := new(reqCreate)
	//if err := c.Bind(req); err != nil {
	//	return err
	//}

	reqBody := new(reqCreateBody)
	if err := c.Bind(reqBody); err != nil {
		return err
	}

	game, err := h.GameService.GetByID(c.Request().Context(), c.Param("game_id"))
	if err != nil {
		return err
	}

	media, err := h.GameMediaService.Create(c.Request().Context(), &service.CreateGameMediaData{
		Game: game,
		Type: game_media.NewTypeByString(reqBody.Type),
	})

	if err != nil {
		return err
	}

	return response.New(c, h.view(media))
}
