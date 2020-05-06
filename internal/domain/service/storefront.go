package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/storefront_service.go -package=mocks github.com/qilin/crm-api/internal/domain/service StorefrontService
type StorefrontService interface {
	Create(ctx context.Context, data *entity.Storefront) (*entity.Storefront, error)
	Update(ctx context.Context, data *entity.Storefront) (*entity.Storefront, error)
	Delete(ctx context.Context, id uint) error
	Activate(ctx context.Context, id uint) error

	GetByID(ctx context.Context, id uint) (*entity.Storefront, error)
	GetAll(ctx context.Context) ([]*entity.Storefront, error)
	FindActive(ctx context.Context) (*entity.Storefront, error)
}
