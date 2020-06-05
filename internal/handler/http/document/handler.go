package document

import (
	"github.com/labstack/echo/v4"
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/pkg/response"
)

// swagger:route POST /documents documents reqDocumentUpsert
//
// Create or update document data
//
// This endpoint will create or update the document and return the document structure
//
//     Responses:
//       200: Document
//       400: HTTPError
//       404: HTTPError
//       500: HTTPError
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

// swagger:route GET /documents/{id} documents reqDocumentGetByID
//
// Getting a document by ID
//
// This endpoint returns the document structure
//
//     Responses:
//       200: Document
//       400: HTTPError
//       404: HTTPError
//       500: HTTPError
func (h Handler) GetByID(c echo.Context) error {
	var req reqByID
	if err := c.Bind(&req); err != nil {
		return err
	}

	doc, err := h.DocumentService.GetByID(c.Request().Context(), req.ID)
	if err != nil {
		return err
	}

	return response.New(c, h.view(doc))
}

// swagger:route GET /documents documents reqDocumentGetByFilter
//
// Getting a list of documents by filter
//
// This endpoint returns a list of document structures
//
//     Responses:
//       200: DocumentList
//       500: HTTPError
func (h Handler) GetByFilter(c echo.Context) error {
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

// swagger:route POST /documents/{id}/activate documents reqDocumentActivate
//
// Activate the document
//
// This endpoint will activate the document and return the document structure
//
//     Responses:
//       200: Document
//       400: HTTPError
//       404: HTTPError
//       500: HTTPError
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
