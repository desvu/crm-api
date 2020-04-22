package postgres

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/enum/game_publish"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID        uint      `pg:"id"`
	GameID    string    `pg:"game_id,notnull,use_zero"`
	Body      string    `pg:"body,notnull,use_zero"`
	Status    uint8     `pg:"status,notnull,use_zero"`
	CreatedAt time.Time `pg:"created_at"`

	tableName struct{} `pg:"game_store_publish"`
}

func (m model) Convert() *entity.GameStorePublish {
	return &entity.GameStorePublish{
		ID:        m.ID,
		GameID:    m.GameID,
		Body:      m.Body,
		Status:    game_publish.NewStatus(m.Status),
		CreatedAt: m.CreatedAt,
	}
}

func newModel(i *entity.GameStorePublish) *model {
	return &model{
		ID:        i.ID,
		GameID:    i.GameID,
		Body:      i.Body,
		Status:    i.Status.Value(),
		CreatedAt: i.CreatedAt,
	}
}
