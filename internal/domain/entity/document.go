package entity

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/enum/document"
	"github.com/qilin/crm-api/internal/domain/enum/language"
)

type Document struct {
	ID          uint
	Title       string
	Text        string
	Reason      string
	Type        document.Type
	Language    language.Language
	Version     string
	CreatedBy   string
	ActivatedBy string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	ActivatedAt *time.Time
}

type DocumentArray []Document

func NewDocumentArray(items []Document) DocumentArray {
	return DocumentArray(items)
}

func (a DocumentArray) IDs() []uint {
	ids := make([]uint, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}

	return ids
}
