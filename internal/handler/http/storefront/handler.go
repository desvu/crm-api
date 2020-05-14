package storefront

import (
	"encoding/json"
	"net/http"

	"github.com/qilin/crm-api/pkg/response"

	"github.com/labstack/echo/v4"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//swagger:parameters deleteStorefronts activateStorefront getStorefront
type requestByID struct {
	// in: path
	// example: 12
	ID uint `param:"id`
}

//swagger:parameters updateStorefronts
type updateRequest struct {
	// in: path
	// example: 12
	ID uint `param:"id`

	// in: body
	Data storefront
}

func (ur *updateRequest) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &ur.Data)
}

//swagger:parameters createStorefronts
type createRequest struct {
	// in: body
	Data storefront
}

// swagger:route GET /storefronts storefronts listStorefronts
//
// Lists all storefront page templates
//
// This endpoint lists all storefront page template
//
//     Responses:
//       default: HTTPError
//       200: StorefrontList
//       400: HTTPError
//       404: HTTPError
//       500: HTTPError
func (h *Handler) List(c echo.Context) error {
	res, err := h.Storefronts.GetAll(c.Request().Context())
	if err != nil {
		return err
	}

	return response.New(c, h.viewList(res))
}

// swagger:route POST /storefronts storefronts createStorefronts
//
// Creates new storefront page template
//
// This endpoint creates new storefront page template
//
//     Responses:
//       default: HTTPError
//       200: Storefront
//       400: HTTPError
//       404: HTTPError
//       500: HTTPError
func (h *Handler) Create(c echo.Context) error {
	var request createRequest
	if err := c.Bind(&request.Data); err != nil {
		return err
	}

	data := &entity.Storefront{
		Name:   request.Data.Name,
		Blocks: request.Data.GetBlocks(),
	}

	res, err := h.Storefronts.Create(c.Request().Context(), data)
	if err != nil {
		return err
	}

	return response.New(c, h.view(res))
}

// swagger:route PUT /storefronts/{id} storefronts updateStorefronts
//
// Updates storefront page template
//
// This endpoint updates storefront page template
//
//     Responses:
//       default: HTTPError
//       200: Storefront
//       400: HTTPError
//       404: HTTPError
//       500: HTTPError
func (h *Handler) Update(ctx echo.Context) error {
	cnt := ctx.Request().Context()

	var request updateRequest
	if err := ctx.Bind(&request); err != nil {
		return err
	}

	data := &entity.Storefront{
		ID:     request.ID,
		Name:   request.Data.Name,
		Blocks: request.Data.GetBlocks(),
	}

	res, err := h.Storefronts.Update(cnt, data)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, h.view(res))
}

// swagger:route DELETE /storefronts/{id} storefronts deleteStorefronts
//
// Removes storefront page template
//
// This endpoint removes storefront page template
//
//     Responses:
//       default: HTTPError
//       204:
//       400: HTTPError
//       404: HTTPError
//       500: HTTPError
func (h *Handler) Delete(ctx echo.Context) error {
	cnt := ctx.Request().Context()

	var req requestByID
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	err := h.Storefronts.Delete(cnt, req.ID)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

// swagger:route GET /storefronts/{id} storefronts getStorefront
//
// Finds storefront page template
//
// This endpoint finds storefront page template
//
//     Responses:
//       default: HTTPError
//       200: Storefront
//       400: HTTPError
//       404: HTTPError
//       500: HTTPError
func (h *Handler) Get(ctx echo.Context) error {
	cnt := ctx.Request().Context()
	var req requestByID
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	res, err := h.Storefronts.GetByID(cnt, req.ID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, h.view(res))
}

// swagger:route POST /storefronts/{id}/activate storefronts activateStorefront
//
// Activates storefront page template
//
// This endpoint activates storefront page template
//
//     Responses:
//       default: HTTPError
//       204:
//       400: HTTPError
//       404: HTTPError
//       500: HTTPError
func (h *Handler) Activate(ctx echo.Context) error {
	cnt := ctx.Request().Context()
	var req requestByID
	if err := ctx.Bind(&req); err != nil {
		return err
	}

	err := h.Storefronts.Activate(cnt, req.ID)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
