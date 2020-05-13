package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/game_revision_genre_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository GameRevisionGenreRepository
type GameRevisionGenreRepository interface {
	Create(ctx context.Context, i *entity.GameRevisionGenre) error
	CreateMultiple(ctx context.Context, items []entity.GameRevisionGenre) error
	Delete(ctx context.Context, i *entity.GameRevisionGenre) error
	DeleteMultiple(ctx context.Context, items []entity.GameRevisionGenre) error

	FindByGenreID(ctx context.Context, genreID uint) ([]entity.GameRevisionGenre, error)
	FindByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.GameRevisionGenre, error)
	FindByGameRevisionIDs(ctx context.Context, gameRevisionIDs []uint) ([]entity.GameRevisionGenre, error)
	FindByGameRevisionIDAndGenreIDs(ctx context.Context, gameRevisionID uint, genreID []uint) ([]entity.GameRevisionGenre, error)
}
