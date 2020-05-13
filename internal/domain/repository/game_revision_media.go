package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type GameRevisionMediaRepository interface {
	Create(ctx context.Context, i *entity.GameRevisionMedia) error
	CreateMultiple(ctx context.Context, items []entity.GameRevisionMedia) error
	Delete(ctx context.Context, i *entity.GameRevisionMedia) error
	DeleteMultiple(ctx context.Context, items []entity.GameRevisionMedia) error

	FindByRevisionID(ctx context.Context, revisionID uint) ([]entity.GameRevisionMedia, error)
	FindByRevisionIDs(ctx context.Context, revisionIDs []uint) ([]entity.GameRevisionMedia, error)
}
