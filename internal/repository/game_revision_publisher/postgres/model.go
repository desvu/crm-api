package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID             uint `pg:"id"`
	GameRevisionID uint `pg:"game_revision_id"`
	PublisherID    uint `pg:"publisher_id"`

	tableName struct{} `pg:"game_revision_publishers"`
}

func (m model) Convert() *entity.GameRevisionPublisher {
	return &entity.GameRevisionPublisher{
		ID:             m.ID,
		GameRevisionID: m.GameRevisionID,
		PublisherID:    m.PublisherID,
	}
}

func newModel(i *entity.GameRevisionPublisher) (*model, error) {
	return &model{
		ID:             i.ID,
		GameRevisionID: i.GameRevisionID,
		PublisherID:    i.PublisherID,
	}, nil
}
