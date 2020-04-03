package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/genre_service.go -package=mocks github.com/qilin/crm-api/internal/domain/service GenreService
type GenreService interface {
	Create(ctx context.Context, data *CreateGenreData) (*entity.Genre, error)
	Update(ctx context.Context, data *UpdateGenreData) (*entity.Genre, error)
	Delete(ctx context.Context, id uint) error

	GetByID(ctx context.Context, id uint) (*entity.Genre, error)
	GetExistByID(ctx context.Context, id uint) (*entity.Genre, error)
	GetByIDs(ctx context.Context, ids []uint) ([]entity.Genre, error)
	GetByGameID(ctx context.Context, gameID uint) ([]entity.Genre, error)

	UpdateGenreForGame(ctx context.Context, game *entity.Game, genreIDs []uint) error
}

type CreateGenreData struct {
	Name string
}

type UpdateGenreData struct {
	ID   uint
	Name string
}
