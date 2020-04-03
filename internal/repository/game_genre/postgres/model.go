package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID      uint `pg:"id"`
	GameID  uint `pg:"game_id"`
	GenreID uint `pg:"genre_id"`

	tableName struct{} `pg:"game_genres"`
}

func (m model) Convert() *entity.GameGenre {
	return &entity.GameGenre{
		ID:      m.ID,
		GameID:  m.GameID,
		GenreID: m.GenreID,
	}
}

func newModel(i *entity.GameGenre) (*model, error) {
	return &model{
		ID:      i.ID,
		GameID:  i.GameID,
		GenreID: i.GenreID,
	}, nil
}
