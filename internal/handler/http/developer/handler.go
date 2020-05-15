package developer

import (
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/pkg/response"

	"github.com/labstack/echo/v4"
)

// swagger:route GET /developers developers listDevelopers
//
// Lists all available developers
//
// This endpoint lists all available developers
//
//     Responses:
//       default: HTTPError
//       200: DeveloperList
//       400: HTTPError
//       404: HTTPError
//       500: HTTPError
func (h *Handler) List(c echo.Context) error {
	data, err := convertGetByFilterRequest(c)
	if err != nil {
		return err
	}

	res, err := h.Developer.GetByFilter(c.Request().Context(), data)
	if err != nil {
		return err
	}

	return response.New(c, h.viewList(res))
}

//swagger:parameters reqGetByFilter
type reqGetByFilter struct {
	// in: query
	// example: 30
	Limit int `query:"limit"`

	// in: query
	// example: 20
	Offset int `query:"offset"`
}

func convertGetByFilterRequest(c echo.Context) (*service.GetByFilterDeveloperData, error) {
	req := new(reqGetByFilter)
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	return &service.GetByFilterDeveloperData{
		Limit:  req.Limit,
		Offset: req.Offset,
	}, nil
}
