package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type FeatureService interface {
	Create(ctx context.Context, data *CreateFeatureData) (*entity.Feature, error)
	Update(ctx context.Context, data *UpdateFeatureData) (*entity.Feature, error)
	Delete(ctx context.Context, id uint) error

	GetByID(ctx context.Context, id uint) (*entity.Feature, error)
	GetExistByID(ctx context.Context, id uint) (*entity.Feature, error)
	GetByIDs(ctx context.Context, ids []uint) ([]entity.Feature, error)
	GetByGameID(ctx context.Context, gameID uint) ([]entity.Feature, error)

	UpdateFeaturesForGame(ctx context.Context, game *entity.Game, featureIDs []uint) error
}

type CreateFeatureData struct {
	Name string
}

type UpdateFeatureData struct {
	ID   uint
	Name string
}
