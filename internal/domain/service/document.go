package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/document"
	"github.com/qilin/crm-api/internal/domain/enum/enum"
	"github.com/qilin/crm-api/internal/domain/enum/language"
)

//go:generate mockgen -destination=../mocks/document_service.go -package=mocks github.com/qilin/crm-api/internal/domain/service DocumentService
type DocumentService interface {
	Create(ctx context.Context, data *CommonDocumentData) (*entity.Document, error)
	Update(ctx context.Context, data *UpdateDocumentData) (*entity.Document, error)
	Upsert(ctx context.Context, data *UpsertDocumentData) (*entity.Document, error)
	Delete(ctx context.Context, id uint) error
	Activate(ctx context.Context, id uint) error

	GetByID(ctx context.Context, id uint) (*entity.Document, error)
	GetExByID(ctx context.Context, id uint) (*entity.Document, error)

	// GetExByFilter returns games founded by filter
	GetExByFilter(ctx context.Context, data *GetByFilterDocumentData) ([]entity.Document, error)
	GetCountByFilter(ctx context.Context, data *GetByFilterDocumentData) (int, error)
}

type CommonDocumentData struct {
	Title     string
	Text      string
	Reason    string
	Type      document.Type
	Language  language.Language
	Version   string
	CreatedBy string
}

func (d CommonDocumentData) Validate() error {
	return nil
}

type UpdateDocumentData struct {
	ID uint
	CommonDocumentData
}

func (d UpdateDocumentData) Validate() error {
	return nil
}

type UpsertDocumentData struct {
	ID *uint
	CommonDocumentData
}

func (d UpsertDocumentData) Validate() error {
	return nil
}

type GetByFilterDocumentData struct {
	OnlyActivated bool
	OrderType     enum.SortOrderType
	Limit         int
	Offset        int
}
