package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type PublisherRepository interface {
	Create(ctx context.Context, i *entity.Publisher) error
	Update(ctx context.Context, i *entity.Publisher) error
	Delete(ctx context.Context, i *entity.Publisher) error

	FindByID(ctx context.Context, id uint) (*entity.Publisher, error)
	FindByIDs(ctx context.Context, ids []uint) ([]entity.Publisher, error)
}
