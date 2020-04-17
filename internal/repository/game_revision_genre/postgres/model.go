package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID             uint `pg:"id"`
	GameRevisionID uint `pg:"game_revision_id"`
	GenreID        uint `pg:"genre_id"`

	tableName struct{} `pg:"game_revision_genres"`
}

func (m model) Convert() *entity.GameRevisionGenre {
	return &entity.GameRevisionGenre{
		ID:             m.ID,
		GameRevisionID: m.GameRevisionID,
		GenreID:        m.GenreID,
	}
}

func newModel(i *entity.GameRevisionGenre) (*model, error) {
	return &model{
		ID:             i.ID,
		GameRevisionID: i.GameRevisionID,
		GenreID:        i.GenreID,
	}, nil
}
