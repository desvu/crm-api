package entity

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/enum/game_publish"
)

type GameStorePublish struct {
	ID        uint
	GameID    string
	Body      string
	Status    game_publish.Status
	CreatedAt time.Time
}
