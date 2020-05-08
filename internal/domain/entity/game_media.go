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
