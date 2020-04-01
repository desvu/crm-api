package entity

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/enum/game"
)

type Game struct {
	ID          uint
	Title       string
	Summary     string
	Description string
	License     string
	Ranking     string
	Type        game.Type
	Platforms   game.PlatformArray
	ReleaseDate time.Time
}
