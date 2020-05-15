package genre

import (
	"github.com/qilin/crm-api/pkg/response"

	"github.com/labstack/echo/v4"
)

// swagger:route GET /genres genres listGenres
//
// Lists all available genres
//
// This endpoint lists all available genres
//
//     Responses:
//       default: HTTPError
//       200: GenreList
//       400: HTTPError
//       404: HTTPError
//       500: HTTPError
func (h *Handler) List(c echo.Context) error {
	res, err := h.Genre.GetAll(c.Request().Context())
	if err != nil {
		return err
	}

	return response.New(c, h.viewList(res))
}
