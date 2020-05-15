package feature

import (
	"github.com/qilin/crm-api/pkg/response"

	"github.com/labstack/echo/v4"
)

// swagger:route GET /features features listFeatures
//
// Lists all available features
//
// This endpoint lists all available features
//
//     Responses:
//       default: HTTPError
//       200: FeatureList
//       500: HTTPError
func (h *Handler) List(c echo.Context) error {
	res, err := h.Feature.GetAll(c.Request().Context())
	if err != nil {
		return err
	}

	return response.New(c, h.viewList(res))
}
