package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/game_revision_feature_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository GameRevisionFeatureRepository
type GameRevisionFeatureRepository interface {
	Create(ctx context.Context, i *entity.GameRevisionFeature) error
	CreateMultiple(ctx context.Context, items []entity.GameRevisionFeature) error
	Delete(ctx context.Context, i *entity.GameRevisionFeature) error
	DeleteMultiple(ctx context.Context, items []entity.GameRevisionFeature) error

	FindByFeatureID(ctx context.Context, featureID uint) ([]entity.GameRevisionFeature, error)
	FindByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.GameRevisionFeature, error)
	FindByGameRevisionIDAndFeatureIDs(ctx context.Context, gameRevisionID uint, featureIDs []uint) ([]entity.GameRevisionFeature, error)
}
