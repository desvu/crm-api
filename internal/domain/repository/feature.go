package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/feature_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository FeatureRepository
type FeatureRepository interface {
	Create(ctx context.Context, i *entity.Feature) error
	Update(ctx context.Context, i *entity.Feature) error
	Delete(ctx context.Context, i *entity.Feature) error

	FindAll(ctx context.Context) ([]entity.Feature, error)
	FindByID(ctx context.Context, id uint) (*entity.Feature, error)
	FindByIDs(ctx context.Context, ids []uint) ([]entity.Feature, error)
}
