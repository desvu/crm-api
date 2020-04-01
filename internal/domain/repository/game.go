package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type IGameRepository interface {
	Create(ctx context.Context, i *entity.Game) error
	//Update(ctx context.Context, i *entity.Game) error
	//Delete(ctx context.Context, i *entity.Game) error

	FindByID(ctx context.Context, id uint) (*entity.Game, error)
}
