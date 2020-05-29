package document

import (
	"github.com/labstack/echo/v4"
	"github.com/qilin/crm-api/internal/domain/enum/document"
	"github.com/qilin/crm-api/internal/domain/enum/language"
	"github.com/qilin/crm-api/internal/domain/service"
)

//swagger:parameters reqDocumentUpsert
type reqUpsert struct {
	ID       *uint  `json:"id"`
	Text     string `json:"text"`
	Reason   string `json:"reason"`
	Type     string `json:"type"`
	Language string `json:"language"`
	Version  string `json:"version"`
}

func convertUpsertRequest(c echo.Context) (*service.UpsertDocumentData, error) {
	req := new(reqUpsert)
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	data := &service.UpsertDocumentData{
		ID: req.ID,
		CommonDocumentData: service.CommonDocumentData{
			Text:     req.Text,
			Reason:   req.Reason,
			Type:     document.NewTypeByString(req.Type),
			Language: language.NewLanguage(req.Language),
			Version:  req.Version,
		},
	}

	return data, nil
}

//swagger:parameters reqDocumentGetByFilter
type reqGetByFilter struct {
	// in: query
	// example: true
	OnlyActivated bool `query:"only_activated"`

	// in: query
	// example: 30
	Limit int `query:"limit"`

	// in: query
	// example: 20
	Offset int `query:"offset"`
}

func convertGetByFilterRequest(c echo.Context) (*service.GetByFilterDocumentData, error) {
	req := new(reqGetByFilter)
	if err := c.Bind(req); err != nil {
		return nil, err
	}

	return &service.GetByFilterDocumentData{
		OnlyActivated: req.OnlyActivated,
		Limit:         req.Limit,
		Offset:        req.Offset,
	}, nil
}

//swagger:parameters reqDocumentGetByID reqDocumentActivate
type reqByID struct {
	// in: path
	// required: true
	// example: 10
	ID uint `param:"id"`
}
