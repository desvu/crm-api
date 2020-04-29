package service

import (
	"context"
	"time"

	"github.com/qilin/crm-api/internal/domain/enum/game"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/game_service.go -package=mocks github.com/qilin/crm-api/internal/domain/service GameService
type GameService interface {
	Create(ctx context.Context, data *CreateGameData) (*entity.GameEx, error)
	Update(ctx context.Context, data *UpdateGameData) (*entity.GameEx, error)
	Delete(ctx context.Context, id string) error
	Publish(ctx context.Context, id string) error

	GetByID(ctx context.Context, id string) (*entity.Game, error)
	GetExByID(ctx context.Context, id string) (*entity.GameEx, error)
	GetExByIDAndRevisionID(ctx context.Context, id string, revisionID uint) (*entity.GameEx, error)
	// GetExBySlig returns last published game by slug
	GetExBySlug(ctx context.Context, slug string) (*entity.GameEx, error)
}

type CreateGameData struct {
	Title       string
	Slug        string
	Type        game.Type
	Summary     *string
	Description *string
	License     *string
	Platforms   *game.PlatformArray
	ReleaseDate *time.Time

	Tags       *[]uint
	Developers *[]uint
	Publishers *[]uint
	Features   *[]uint
	Genres     *[]uint
}

type UpdateGameData struct {
	ID         string
	Title      *string
	Slug       *string
	Tags       *[]uint
	Developers *[]uint
	Publishers *[]uint
	Features   *[]uint
	Genres     *[]uint
}
