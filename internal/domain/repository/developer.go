package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/developer_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository DeveloperRepository
type DeveloperRepository interface {
	Create(ctx context.Context, i *entity.Developer) error
	Update(ctx context.Context, i *entity.Developer) error
	Delete(ctx context.Context, i *entity.Developer) error

	FindByID(ctx context.Context, id uint) (*entity.Developer, error)
	FindByIDs(ctx context.Context, ids []uint) ([]entity.Developer, error)
	FindByFilter(ctx context.Context, data *FindByFilterDeveloperData) ([]entity.Developer, error)
}

type FindByFilterDeveloperData struct {
	Limit  int
	Offset int
}
