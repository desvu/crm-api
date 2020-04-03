package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID          uint `pg:"id"`
	GameID      uint `pg:"game_id"`
	DeveloperID uint `pg:"developer_id"`

	tableName struct{} `pg:"game_developers"`
}

func (m model) Convert() *entity.GameDeveloper {
	return &entity.GameDeveloper{
		ID:          m.ID,
		GameID:      m.GameID,
		DeveloperID: m.DeveloperID,
	}
}

func newModel(i *entity.GameDeveloper) (*model, error) {
	return &model{
		ID:          i.ID,
		GameID:      i.GameID,
		DeveloperID: i.DeveloperID,
	}, nil
}
