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
	Trailer            string
	PlayTime           uint
	ReleaseDate        time.Time
	PublishedAt        *time.Time
	SystemRequirements []SystemRequirements
	SocialLinks        []SocialLink

	Status    game_revision.Status
	Platforms game.PlatformArray
}

type GameRevisionArray []GameRevision

func NewGameRevisionArray(items []GameRevision) GameRevisionArray {
	return GameRevisionArray(items)
}

func (a GameRevisionArray) IDs() []uint {
	ids := make([]uint, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}

	return ids
}

func (a GameRevisionArray) GameIDs() []string {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].GameID
	}

	return ids
}
