package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID     uint `pg:"id"`
	GameID uint `pg:"game_id"`
	TagID  uint `pg:"tag_id"`

	tableName struct{} `pg:"game_tags"`
}

func (m model) Convert() *entity.GameTag {
	return &entity.GameTag{
		ID:     m.ID,
		GameID: m.GameID,
		TagID:  m.TagID,
	}
}

func newModel(i *entity.GameTag) (*model, error) {
	return &model{
		ID:     i.ID,
		GameID: i.GameID,
		TagID:  i.TagID,
	}, nil
}
