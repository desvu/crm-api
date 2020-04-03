package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/genre_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository GenreRepository
type GenreRepository interface {
	Create(ctx context.Context, i *entity.Genre) error
	Update(ctx context.Context, i *entity.Genre) error
	Delete(ctx context.Context, i *entity.Genre) error

	FindByID(ctx context.Context, id uint) (*entity.Genre, error)
	FindByIDs(ctx context.Context, ids []uint) ([]entity.Genre, error)
}
