package entity

import "github.com/qilin/crm-api/internal/domain/enum/game"

type Feature struct {
	ID   uint
	Name string
	Icon game.Icon
}
