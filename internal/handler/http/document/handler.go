package document

import (
	"github.com/labstack/echo/v4"
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/pkg/response"
)

func (h Handler) Upsert(c echo.Context) error {
	data, err := convertUpsertRequest(c)
	if err != nil {
		return err
	}

	doc, err := h.DocumentService.Upsert(c.Request().Context(), data)
	if err != nil {
		return err
	}

	return response.New(c, h.view(doc))
}

func (h Handler) List(c echo.Context) error {
	data, err := convertGetByFilterRequest(c)
	if err != nil {
		return err
	}

	docs, err := h.DocumentService.GetExByFilter(c.Request().Context(), &service.GetByFilterDocumentData{
		OnlyActivated: data.OnlyActivated,
		Limit:         data.Limit,
		Offset:        data.Offset,
	})

	totalCount, err := h.DocumentService.GetCountByFilter(c.Request().Context(), data)
	if err != nil {
		return err
	}

	return response.New(c, h.viewArray(docs, totalCount))
}

func (h Handler) Activate(c echo.Context) error {
	var req reqByID
	if err := c.Bind(&req); err != nil {
		return err
	}

	err := h.DocumentService.Activate(c.Request().Context(), req.ID)
	if err != nil {
		return err
	}

	doc, err := h.DocumentService.GetByID(c.Request().Context(), req.ID)
	if err != nil {
		return err
	}

	return response.New(c, h.view(doc))
}
