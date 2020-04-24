package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type GameStorePublishRepository interface {
	Create(ctx context.Context, i *entity.GameStorePublish) error
	Update(ctx context.Context, i *entity.GameStorePublish) error
	Delete(ctx context.Context, i *entity.GameStorePublish) error

	FindByID(ctx context.Context, id uint) (*entity.GameStorePublish, error)
	FindByGameID(ctx context.Context, gameID string) ([]entity.GameStorePublish, error)
}
