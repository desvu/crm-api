package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID        uint `pg:"id"`
	GameID    uint `pg:"game_id"`
	FeatureID uint `pg:"developer_id"`

	tableName struct{} `pg:"game_features"`
}

func (m model) Convert() *entity.GameFeature {
	return &entity.GameFeature{
		ID:        m.ID,
		GameID:    m.GameID,
		FeatureID: m.FeatureID,
	}
}

func newModel(i *entity.GameFeature) (*model, error) {
	return &model{
		ID:        i.ID,
		GameID:    i.GameID,
		FeatureID: i.FeatureID,
	}, nil
}
