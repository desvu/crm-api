package entity

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/enum/game"
	"github.com/qilin/crm-api/internal/domain/enum/game_revision"
)

type GameRevision struct {
	ID                 uint
	GameID             string
	Summary            string
	Description        string
	License            string
	ReleaseDate        time.Time
	PublishedAt        *time.Time
	SystemRequirements []SystemRequirements
	SocialLinks        []SocialLink

	Status    game_revision.Status
	Platforms game.PlatformArray
}
