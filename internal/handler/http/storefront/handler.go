package storefront

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/qilin/crm-api/internal/domain/entity"
	derrors "github.com/qilin/crm-api/internal/domain/errors"
	perrors "github.com/qilin/crm-api/pkg/errors"
)

type requestData struct {
	Name   string         `json:"name"`
	Blocks []entity.Block `json:"blocks"`
}

// swagger:route GET /storefronts storefronts listStorefronts
//
// Lists all storefront page templates
//
// This endpoint lists all storefront page template
//
//     Responses:
//       200: StorefrontList
func (h *Handler) List(ctx echo.Context) error {
	cnt := ctx.Request().Context()

	res, err := h.Storefronts.GetAll(cnt)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, h.viewList(res))
}

// swagger:route POST /storefronts storefronts createStorefronts
//
// Creates new storefront page template
//
// This endpoint creates new storefront page template
//
//     Responses:
//       200: Storefront
func (h *Handler) Create(ctx echo.Context) error {
	cnt := ctx.Request().Context()

	var request requestData
	if err := ctx.Bind(&request); err != nil {
		return err
	}

	data := &entity.Storefront{
		Name:   request.Name,
		Blocks: request.Blocks,
	}

	res, err := h.Storefronts.Create(cnt, data)
	if err != nil {
		var x perrors.Error
		if errors.As(err, &x) {
			if x.Type != perrors.ErrInternal {
				return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
			}
		}
		return err
	}

	return ctx.JSON(http.StatusOK, h.view(res))
}

// swagger:route PUT /storefronts/:id storefronts updateStorefronts
//
// Updates storefront page template
//
// This endpoint updates storefront page template
//
//     Responses:
//       200: Storefront
func (h *Handler) Update(ctx echo.Context) error {
	cnt := ctx.Request().Context()
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return h.NotFound(ctx, err)
	}

	var request requestData
	if err := ctx.Bind(&request); err != nil {
		return err
	}

	data := &entity.Storefront{
		ID:     uint(id),
		Name:   request.Name,
		Blocks: request.Blocks,
	}

	res, err := h.Storefronts.Update(cnt, data)
	if err != nil {
		var x perrors.Error
		if errors.As(err, &x) {
			if x.Type != perrors.ErrInternal {
				return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
			}
		}
		return err
	}

	return ctx.JSON(http.StatusOK, h.view(res))
}

// swagger:route DELETE /storefronts/:id storefronts deleteStorefronts
//
// Removes storefront page template
//
// This endpoint removes storefront page template
//
//     Responses:
//       204:
func (h *Handler) Delete(ctx echo.Context) error {
	cnt := ctx.Request().Context()
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return h.NotFound(ctx, err)
	}

	err = h.Storefronts.Delete(cnt, uint(id))
	if err != nil {
		if errors.Is(err, derrors.StoreFrontNotFound) {
			return h.NotFound(ctx, err)
		}
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

// swagger:route GET /storefronts/:id storefronts getStorefront
//
// Finds storefront page template
//
// This endpoint finds storefront page template
//
//     Responses:
//       200: Storefront
func (h *Handler) Get(ctx echo.Context) error {
	cnt := ctx.Request().Context()
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return h.NotFound(ctx, err)
	}

	res, err := h.Storefronts.GetByID(cnt, uint(id))
	if err != nil {
		if errors.Is(err, derrors.StoreFrontNotFound) {
			return h.NotFound(ctx, err)
		}
		return err
	}

	return ctx.JSON(http.StatusOK, h.view(res))
}

// swagger:route POST /storefronts/:id/activate storefronts activateStorefront
//
// Activates storefront page template
//
// This endpoint activates storefront page template
//
//     Responses:
//       204:
func (h *Handler) Activate(ctx echo.Context) error {
	cnt := ctx.Request().Context()
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return h.NotFound(ctx, err)
	}

	err = h.Storefronts.Activate(cnt, uint(id))
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *Handler) NotFound(ctx echo.Context, err error) error {
	// ignore error
	return ctx.JSON(http.StatusNotFound, echo.Map{}) // TODO struct
}
