package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type DocumentRepository interface {
	Create(ctx context.Context, i *entity.Document) error
	Update(ctx context.Context, i *entity.Document) error
	Delete(ctx context.Context, i *entity.Document) error

	FindByID(ctx context.Context, id uint) (*entity.Document, error)
	FindByIDs(ctx context.Context, ids []uint) ([]entity.Document, error)
	FindByFilter(ctx context.Context, filter *FindByFilterDocumentData) ([]entity.Document, error)
	CountByFilter(ctx context.Context, filter *FindByFilterDocumentData) (int, error)
}

type FindByFilterDocumentData struct {
	OnlyActivated bool
	Limit         int
	Offset        int
}
