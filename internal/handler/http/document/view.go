package document

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//swagger:model Document
type doc struct {
	// example: 100
	ID uint `json:"id"`
	// example: Privacy Policy v1.2
	Title string `json:"title"`
	// example: Privacy Policy Description
	Text string `json:"text"`
	// example: Basic Privacy Policy
	Reason string `json:"reason"`
	// example: privacyPolicy
	Type string `json:"type"`
	// example: en
	Language string `json:"language"`
	// example: v1.2
	Version string `json:"version"`
	//
	CreatedBy string `json:"created_by"`
	//
	ActivatedBy string `json:"activated_by"`
	//
	CreatedAt time.Time `json:"created_at"`
	//
	UpdatedAt *time.Time `json:"updated_at"`
	//
	ActivatedAt *time.Time `json:"activated_at"`
}

func (h Handler) view(i *entity.Document) doc {
	return doc{
		ID:          i.ID,
		Title:       i.Title,
		Text:        i.Text,
		Reason:      i.Reason,
		Type:        i.Type.String(),
		Language:    string(i.Language),
		Version:     i.Version,
		CreatedBy:   i.CreatedBy,
		ActivatedBy: i.ActivatedBy,
		CreatedAt:   i.CreatedAt,
		UpdatedAt:   i.UpdatedAt,
		ActivatedAt: i.ActivatedAt,
	}
}

type pagination struct {
	Total int `json:"total"`
}

//swagger:model DocumentList
type documentList struct {
	Documents  []doc      `json:"documents"`
	Pagination pagination `json:"pagination"`
}

func (h Handler) viewArray(items []entity.Document, totalCount int) documentList {
	var docs = make([]doc, len(items))
	for i := range items {
		docs[i] = h.view(&items[i])
	}

	return documentList{
		Documents: docs,
		Pagination: pagination{
			Total: totalCount,
		},
	}
}
