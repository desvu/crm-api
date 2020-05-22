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
//	<p>
//		Поддерживается загрузка только png изображение, которые будут перекодированны
//		в jpg изображения. Размер изображения будет изменен до минимального необходимого
// 		разрешения.
//  </p>
//
//	<p>
// 		В зависимости от типа изображения, к нему выставляются определенные тробования:
//		<table>
//  		<tr>
//    			<th>Тип</th>
//				<th>Необходимое соотношение сторон</th>
//				<th>Минимальное разрешение</th>
//  		</tr>
//  		<tr>
//    			<td>wideSlider</td>
//				<td>16:9</td>
//				<td>1064 * 560</td>
//  		</tr>
//  		<tr>
//    			<td>vertical</td>
//				<td>3:4</td>
//				<td>200 * 266</td>
//  		</tr>
//  		<tr>
//    			<td>horizontal</td>
//				<td>16:9</td>
//				<td>524 * 294</td>
//  		</tr>
//  		<tr>
//    			<td>horizontalSmall</td>
//				<td>16:9</td>
//				<td>254 * 142</td>
//  		</tr>
//  		<tr>
//    			<td>largeSingle</td>
//				<td>16:9</td>
//				<td>744 * 410</td>
//  		</tr>
//  		<tr>
//    			<td>catalog</td>
//				<td>16:9</td>
//				<td>88 * 50</td>
//  		</tr>
//  		<tr>
//    			<td>screenshot</td>
//				<td>16:9</td>
//				<td>1064 * 562</td>
//  		</tr>
//  		<tr>
//    			<td>description</td>
//				<td>-</td>
//				<td>-</td>
//  		</tr>
//		</table>
//	</p>
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
// Creates game media images
//
//	Загрузка изображения происходит в два этапа:
//	<ol>
//		<li>
//			Создание сущности game media, при котором указывается типа изображения. При создании
//			резервируется имя изображения в хранилище и генерируется уникальный идентификатор. При
//			создании, поле is_uploaded имеет значение false. Game media возможно прикрепить к игре,
// 			только после загрузки изображения, после которого поле is_uploaded примет значение true.
//		</li>
//		<li>
//			Загрузка изображения. На этапе загрузки изображениея происходит проверка изображение на
// 			необходимые размеры и типы. После загрузки изображениея, его можно прикрепить к игре.
//		</li>
//	</ol>
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
