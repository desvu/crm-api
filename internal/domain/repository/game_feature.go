package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/game_feature_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository GameFeatureRepository
type GameFeatureRepository interface {
	Create(ctx context.Context, i *entity.GameFeature) error
	CreateMultiple(ctx context.Context, items []entity.GameFeature) error
	Delete(ctx context.Context, i *entity.GameFeature) error
	DeleteMultiple(ctx context.Context, items []entity.GameFeature) error

	FindByFeatureID(ctx context.Context, featureID uint) ([]entity.GameFeature, error)
	FindByGameID(ctx context.Context, gameID uint) ([]entity.GameFeature, error)
	FindByGameIDAndFeatureIDs(ctx context.Context, gameID uint, featureIDs []uint) ([]entity.GameFeature, error)
}
