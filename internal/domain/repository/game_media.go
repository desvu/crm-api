package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type GameMediaRepository interface {
	Create(ctx context.Context, i *entity.GameMedia) error
	Delete(ctx context.Context, i *entity.GameMedia) error

	FindByID(ctx context.Context, id uint) (*entity.GameMedia, error)
	FindByIDs(ctx context.Context, ids []uint) ([]entity.GameMedia, error)
}
