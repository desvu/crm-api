package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type StoreFrontRepository interface {
	Create(ctx context.Context, i *entity.Storefront) error
	Update(ctx context.Context, i *entity.Storefront) error
	Delete(ctx context.Context, id uint) error
	Activate(ctx context.Context, id, version uint) error

	GetByID(ctx context.Context, id uint) (*entity.Storefront, error)
	GetByIDAndVersion(ctx context.Context, id, version uint) (*entity.Storefront, error)
	GetAll(ctx context.Context) ([]*entity.Storefront, error)

	FindActive(ctx context.Context) (*entity.Storefront, error)
}
