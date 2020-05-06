package http

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/qilin/crm-api/internal/domain/entity"
	derrors "github.com/qilin/crm-api/internal/domain/errors"
	"github.com/qilin/crm-api/internal/domain/service"
	perrors "github.com/qilin/crm-api/pkg/errors"
	"go.uber.org/fx"
)

type StorefrontParams struct {
	fx.In

	Storefronts service.StorefrontService
}

type StorefrontHandler struct {
	Storefronts service.StorefrontService
	// px StorefrontParams
}

func NewStorefrontHandler(p StorefrontParams) *StorefrontHandler {
	return &StorefrontHandler{p.Storefronts}
}

func (h *StorefrontHandler) List(ctx echo.Context) error {
	cnt := ctx.Request().Context()

	res, err := h.Storefronts.GetAll(cnt)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, h.convList(res))
}

func (h *StorefrontHandler) Create(ctx echo.Context) error {
	cnt := ctx.Request().Context()

	var request StorefrontRequest
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

	return ctx.JSON(http.StatusOK, h.conv(res))
}

func (h *StorefrontHandler) Update(ctx echo.Context) error {
	cnt := ctx.Request().Context()
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return h.NotFound(ctx, err)
	}

	var request StorefrontRequest
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

	return ctx.JSON(http.StatusOK, h.conv(res))
}

func (h *StorefrontHandler) Delete(ctx echo.Context) error {
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

func (h *StorefrontHandler) Get(ctx echo.Context) error {
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

	return ctx.JSON(http.StatusOK, h.conv(res))
}

func (h *StorefrontHandler) Activate(ctx echo.Context) error {
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

func (h *StorefrontHandler) conv(sf *entity.Storefront) *StorefrontFull {
	return &StorefrontFull{
		Storefront: *h.convBasic(sf),
		Blocks:     sf.Blocks,
	}
}

func (h *StorefrontHandler) convBasic(sf *entity.Storefront) *Storefront {
	return &Storefront{
		ID:        strconv.FormatUint(uint64(sf.ID), 10),
		Name:      sf.Name,
		IsActive:  sf.IsActive,
		Version:   sf.Version,
		CreatedAt: sf.CreatedAt,
		UpdatedAt: sf.UpdatedAt,
	}
}

func (h *StorefrontHandler) convList(sf []*entity.Storefront) []*Storefront {
	var res = make([]*Storefront, len(sf))
	for i := range sf {
		res[i] = h.convBasic(sf[i])
	}
	return res
}

func (h *StorefrontHandler) NotFound(ctx echo.Context, err error) error {
	// ignore error
	return ctx.JSON(http.StatusNotFound, echo.Map{}) // TODO struct
}

type StorefrontRequest struct {
	Name   string         `json:"name"`
	Blocks []entity.Block `json:"blocks"`
}

type Storefront struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Version   uint      `json:"version"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StorefrontFull struct {
	Storefront
	Blocks []entity.Block `json:"blocks"`
}
