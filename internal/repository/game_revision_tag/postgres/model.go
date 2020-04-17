package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID             uint `pg:"id"`
	GameRevisionID uint `pg:"game_revision_id"`
	TagID          uint `pg:"tag_id"`

	tableName struct{} `pg:"game_revision_tags"`
}

func (m model) Convert() *entity.GameRevisionTag {
	return &entity.GameRevisionTag{
		ID:             m.ID,
		GameRevisionID: m.GameRevisionID,
		TagID:          m.TagID,
	}
}

func newModel(i *entity.GameRevisionTag) (*model, error) {
	return &model{
		ID:             i.ID,
		GameRevisionID: i.GameRevisionID,
		TagID:          i.TagID,
	}, nil
}
