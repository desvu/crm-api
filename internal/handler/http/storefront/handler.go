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

func (h *Handler) List(ctx echo.Context) error {
	cnt := ctx.Request().Context()

	res, err := h.Storefronts.GetAll(cnt)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, h.viewList(res))
}

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
