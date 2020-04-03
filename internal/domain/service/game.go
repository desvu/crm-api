package service

import (
	"context"
	"time"

	"github.com/qilin/crm-api/internal/domain/enum/game"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/game_service.go -package=mocks github.com/qilin/crm-api/internal/domain/service GameService
type GameService interface {
	Create(ctx context.Context, data *CreateGameData) (*entity.Game, error)
	Update(ctx context.Context, data *UpdateGameData) (*entity.Game, error)
	UpdateEx(ctx context.Context, data *UpdateGameExData) (*entity.GameEx, error)
	Delete(ctx context.Context, id uint) error
	Publish(ctx context.Context, id uint) error

	GetByID(ctx context.Context, id uint) (*entity.Game, error)
	GetExistByID(ctx context.Context, id uint) (*entity.Game, error)
}

type CreateGameData struct {
	Title       string
	Summary     string
	Description string
	License     string
	Ranking     string
	Type        game.Type
	Platforms   game.PlatformArray
	ReleaseDate time.Time
}

type UpdateGameData struct {
	ID    uint
	Title string
}

type UpdateGameExData struct {
	UpdateGameData
	Tags       *[]uint
	Developers *[]uint
	Publishers *[]uint
	Features   *[]uint
	Genres     *[]uint
}
