package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID          uint `pg:"id"`
	RevisionID  uint `pg:"revision_id,notnull,use_zero"`
	GameMediaID uint `pg:"game_media_id,notnull,use_zero"`

	tableName struct{} `pg:"game_revision_media"`
}

func (m model) Convert() *entity.GameRevisionMedia {
	return &entity.GameRevisionMedia{
		ID:         m.ID,
		RevisionID: m.RevisionID,
		MediaID:    m.GameMediaID,
	}
}

func newModel(i *entity.GameRevisionMedia) *model {
	return &model{
		ID:          i.ID,
		RevisionID:  i.RevisionID,
		GameMediaID: i.MediaID,
	}
}
