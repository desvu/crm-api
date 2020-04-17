package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID             uint `pg:"id"`
	GameRevisionID uint `pg:"game_revision_id"`
	FeatureID      uint `pg:"feature_id"`

	tableName struct{} `pg:"game_revision_features"`
}

func (m model) Convert() *entity.GameRevisionFeature {
	return &entity.GameRevisionFeature{
		ID:             m.ID,
		GameRevisionID: m.GameRevisionID,
		FeatureID:      m.FeatureID,
	}
}

func newModel(i *entity.GameRevisionFeature) (*model, error) {
	return &model{
		ID:             i.ID,
		GameRevisionID: i.GameRevisionID,
		FeatureID:      i.FeatureID,
	}, nil
}
