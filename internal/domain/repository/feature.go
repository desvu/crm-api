package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type FeatureRepository interface {
	Create(ctx context.Context, i *entity.Feature) error
	Update(ctx context.Context, i *entity.Feature) error
	Delete(ctx context.Context, i *entity.Feature) error

	FindByID(ctx context.Context, id uint) (*entity.Feature, error)
	FindByIDs(ctx context.Context, ids []uint) ([]entity.Feature, error)
}
