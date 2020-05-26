package entity

import "github.com/qilin/crm-api/internal/domain/enum/game"

type Game struct {
	ID    string
	Title string
	Slug  string
	Type  game.Type
}

type GameArray []Game

func NewGameArray(items []Game) GameArray {
	return GameArray(items)
}

func (a GameArray) IDs() []string {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}

	return ids
}
