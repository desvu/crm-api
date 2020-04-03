package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID          uint `pg:"id"`
	GameID      uint `pg:"game_id"`
	PublisherID uint `pg:"publisher_id"`

	tableName struct{} `pg:"game_publishers"`
}

func (m model) Convert() *entity.GamePublisher {
	return &entity.GamePublisher{
		ID:          m.ID,
		GameID:      m.GameID,
		PublisherID: m.PublisherID,
	}
}

func newModel(i *entity.GamePublisher) (*model, error) {
	return &model{
		ID:          i.ID,
		GameID:      i.GameID,
		PublisherID: i.PublisherID,
	}, nil
}
