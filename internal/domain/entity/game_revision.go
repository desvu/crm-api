package entity

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/enum/game"
	"github.com/qilin/crm-api/internal/domain/enum/game_revision"
)

type GameRevision struct {
	ID          uint
	GameID      string
	Summary     string
	Description string
	Slug        string
	License     string
	ReleaseDate time.Time
	PublishedAt *time.Time

	Status    game_revision.Status
	Platforms game.PlatformArray
}
