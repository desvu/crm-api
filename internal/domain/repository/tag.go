package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/tag_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository TagRepository
type TagRepository interface {
	Create(ctx context.Context, i *entity.Tag) error
	Update(ctx context.Context, i *entity.Tag) error
	Delete(ctx context.Context, i *entity.Tag) error

	FindAll(ctx context.Context) ([]entity.Tag, error)
	FindByID(ctx context.Context, id uint) (*entity.Tag, error)
	FindByIDs(ctx context.Context, ids []uint) ([]entity.Tag, error)
}
