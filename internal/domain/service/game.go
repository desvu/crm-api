package service

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game"
)

//go:generate mockgen -destination=../mocks/game_service.go -package=mocks github.com/qilin/crm-api/internal/domain/service GameService
type GameService interface {
	Create(ctx context.Context, data *CreateGameData) (*entity.GameEx, error)
	Update(ctx context.Context, data *UpdateGameData) (*entity.GameEx, error)
	Upsert(ctx context.Context, data *UpsertGameData) (*entity.GameEx, error)
	Delete(ctx context.Context, id string) error
	Publish(ctx context.Context, id string) error

	GetByID(ctx context.Context, id string) (*entity.Game, error)
	GetExByID(ctx context.Context, id string) (*entity.GameEx, error)
	GetExByIDAndRevisionID(ctx context.Context, id string, revisionID uint) (*entity.GameEx, error)

	// GetExLastPublishedByID returns last published game by id
	GetExLastPublishedByID(ctx context.Context, id string) (*entity.GameEx, error)

	// GetExBySlug returns last published game by slug
	GetExBySlug(ctx context.Context, slug string) (*entity.GameEx, error)
}

type CommonGameData struct {
	Summary     *string
	Description *string
	License     *string
	Tags        *[]uint
	Developers  *[]uint
	Publishers  *[]uint
	Features    *[]uint
	Genres      *[]uint
	Media       *[]uint

	SystemRequirements *[]SystemRequirements
	Platforms          *game.PlatformArray
	ReleaseDate        *time.Time
}

type UpsertGameData struct {
	ID    *string
	Title *string
	Slug  *string
	Type  *game.Type

	CommonGameData
}

type CreateGameData struct {
	Title string    `validate:"required"`
	Slug  string    `validate:"required"`
	Type  game.Type `validate:"required"`

	CommonGameData
}

func (d CreateGameData) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}

type UpdateGameData struct {
	ID    string
	Title *string
	Slug  *string
	Type  *game.Type

	CommonGameData
}

type SystemRequirements struct {
	Platform    game.Platform
	Minimal     *RequirementsSet
	Recommended *RequirementsSet
}

type RequirementsSet struct {
	CPU       string
	GPU       string
	DiskSpace uint
	RAM       uint
}
