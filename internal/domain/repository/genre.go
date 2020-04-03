package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type GenreRepository interface {
	Create(ctx context.Context, i *entity.Genre) error
	Update(ctx context.Context, i *entity.Genre) error
	Delete(ctx context.Context, i *entity.Genre) error

	FindByID(ctx context.Context, id uint) (*entity.Genre, error)
	FindByIDs(ctx context.Context, ids []uint) ([]entity.Genre, error)
}
