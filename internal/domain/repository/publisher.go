package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/publisher_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository PublisherRepository
type PublisherRepository interface {
	Create(ctx context.Context, i *entity.Publisher) error
	Update(ctx context.Context, i *entity.Publisher) error
	Delete(ctx context.Context, i *entity.Publisher) error

	FindByID(ctx context.Context, id uint) (*entity.Publisher, error)
	FindByIDs(ctx context.Context, ids []uint) ([]entity.Publisher, error)
	FindByFilter(ctx context.Context, data *FindByFilterPublisherData) ([]entity.Publisher, error)
}

type FindByFilterPublisherData struct {
	Limit  int
	Offset int
}
