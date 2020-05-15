package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game"
)

//go:generate mockgen -destination=../mocks/feature_service.go -package=mocks github.com/qilin/crm-api/internal/domain/service FeatureService
type FeatureService interface {
	Create(ctx context.Context, data *CreateFeatureData) (*entity.Feature, error)
	Update(ctx context.Context, data *UpdateFeatureData) (*entity.Feature, error)
	Delete(ctx context.Context, id uint) error

	GetAll(ctx context.Context) ([]entity.Feature, error)
	GetByID(ctx context.Context, id uint) (*entity.Feature, error)
	GetExistByID(ctx context.Context, id uint) (*entity.Feature, error)
	GetByIDs(ctx context.Context, ids []uint) ([]entity.Feature, error)
	GetByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.Feature, error)

	UpdateFeaturesForGameRevision(ctx context.Context, gameRevision *entity.GameRevision, featureIDs []uint) error
}

type CreateFeatureData struct {
	Name string
	Icon game.Icon
}

type UpdateFeatureData struct {
	ID   uint
	Name string
	Icon game.Icon
}
