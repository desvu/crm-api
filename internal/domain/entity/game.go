package entity

import "github.com/qilin/crm-api/internal/domain/enum/game"

type Game struct {
	ID    string
	Title string
	Slug  string
	Type  game.Type
}

type GameArray struct {
	ids   []string
	items []Game
}

func NewGameArray(items []Game) *GameArray {
	a := &GameArray{
		items: items,
	}
	a.refresh()
	return a
}

func (a *GameArray) refresh() {
	ids := make([]string, len(a.items))

	for i := range a.items {
		ids[i] = a.items[i].ID
	}

	a.ids = ids
}

func (a *GameArray) Size() int {
	return len(a.items)
}

func (a *GameArray) IDs() []string {
	items := make([]string, len(a.items))
	copy(items, a.ids)
	return items
}
