package service

import (
	"context"
	"time"

	"github.com/qilin/crm-api/internal/domain/enum/game"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game_revision"
)

type GameRevisionService interface {
	Update(ctx context.Context, data *UpdateGameRevisionData) (*entity.GameRevisionEx, error)

	GetByID(ctx context.Context, id uint) (*entity.GameRevisionEx, error)
	GetByIDAndGameID(ctx context.Context, id uint, gameID string) (*entity.GameRevisionEx, error)
	GetDraftByGame(ctx context.Context, game *entity.Game) (*entity.GameRevisionEx, error)
	GetLastPublishedByGame(ctx context.Context, game *entity.Game) (*entity.GameRevisionEx, error)
	IsGamesPublished(ctx context.Context, ids ...string) error
}

type UpdateGameRevisionData struct {
	ID          uint
	Summary     *string
	Description *string
	License     *string
	Trailer     *string

	Status             *game_revision.Status
	Platforms          *game.PlatformArray
	ReleaseDate        *time.Time
	SystemRequirements *[]SystemRequirements

	Tags       *[]uint
	Developers *[]uint
	Publishers *[]uint
	Features   *[]uint
	Genres     *[]uint
}
