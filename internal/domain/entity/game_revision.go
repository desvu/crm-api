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
	ReleaseDate        time.Time
	PublishedAt        *time.Time
	SystemRequirements []SystemRequirements
	SocialLinks        []SocialLink

	Status    game_revision.Status
	Platforms game.PlatformArray
}

type GameRevisionArray struct {
	ids     []uint
	gameIDs []string
	items   []GameRevision
}

func NewGameRevisionArray(items []GameRevision) *GameRevisionArray {
	a := &GameRevisionArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *GameRevisionArray) refresh() {
	ids := make([]uint, len(a.items))
	gameIDs := make([]string, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
		gameIDs[i] = a.items[i].GameID
	}

	a.ids = ids
	a.gameIDs = gameIDs
}

func (a *GameRevisionArray) Size() int {
	return len(a.items)
}

func (a *GameRevisionArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}

func (a *GameRevisionArray) GenreIDs() []string {
	items := make([]string, len(a.items))
	copy(items, a.gameIDs)
	return items
}
