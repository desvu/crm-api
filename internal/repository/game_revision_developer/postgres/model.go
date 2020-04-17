package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID             uint `pg:"id"`
	GameRevisionID uint `pg:"game_revision_id"`
	DeveloperID    uint `pg:"developer_id"`

	tableName struct{} `pg:"game_revision_developers"`
}

func (m model) Convert() *entity.GameRevisionDeveloper {
	return &entity.GameRevisionDeveloper{
		ID:             m.ID,
		GameRevisionID: m.GameRevisionID,
		DeveloperID:    m.DeveloperID,
	}
}

func newModel(i *entity.GameRevisionDeveloper) (*model, error) {
	return &model{
		ID:             i.ID,
		GameRevisionID: i.GameRevisionID,
		DeveloperID:    i.DeveloperID,
	}, nil
}
