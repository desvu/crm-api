package entity

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/enum/game_media"
)

type GameMedia struct {
	ID         uint
	GameID     string
	Type       game_media.Type
	FilePath   string
	IsUploaded bool
	CreatedAt  time.Time
}

type GameMediaArray struct {
	ids   []uint
	items []GameMedia
}

func NewGameMediaArray(items []GameMedia) *GameMediaArray {
	a := &GameMediaArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *GameMediaArray) refresh() {
	ids := make([]uint, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
	}

	a.ids = ids
}

func (a *GameMediaArray) Size() int {
	return len(a.items)
}

func (a *GameMediaArray) IDs() []uint {
	items := make([]uint, len(a.items))
	copy(items, a.ids)
	return items
}
