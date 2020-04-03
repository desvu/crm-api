package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type GameGenreRepository interface {
	Create(ctx context.Context, i *entity.GameGenre) error
	CreateMultiple(ctx context.Context, items []entity.GameGenre) error
	Delete(ctx context.Context, i *entity.GameGenre) error
	DeleteMultiple(ctx context.Context, items []entity.GameGenre) error

	FindByGenreID(ctx context.Context, genreID uint) ([]entity.GameGenre, error)
	FindByGameID(ctx context.Context, gameID uint) ([]entity.GameGenre, error)
	FindByGameIDAndGenreIDs(ctx context.Context, gameID uint, genreID []uint) ([]entity.GameGenre, error)
}
