package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type GameRepository interface {
	Create(ctx context.Context, i *entity.Game) error
	Update(ctx context.Context, i *entity.Game) error
	Delete(ctx context.Context, i *entity.Game) error

	FindByID(ctx context.Context, id uint) (*entity.Game, error)
	FindByIDs(ctx context.Context, ids []uint) ([]entity.Game, error)
}
