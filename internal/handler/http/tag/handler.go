package tag

import (
	"github.com/qilin/crm-api/pkg/response"

	"github.com/labstack/echo/v4"
)

// swagger:route GET /tags tags listTags
//
// Lists all available tags
//
// This endpoint lists all available tags
//
//     Responses:
//       default: HTTPError
//       200: TagList
//       500: HTTPError
func (h *Handler) List(c echo.Context) error {
	res, err := h.Tag.GetAll(c.Request().Context())
	if err != nil {
		return err
	}

	return response.New(c, h.viewList(res))
}
