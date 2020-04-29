package entity

import "github.com/qilin/crm-api/internal/domain/enum/game"

type Game struct {
	ID    string
	Title string
	Slug  string
	Type  game.Type
}
