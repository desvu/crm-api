package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID         uint   `pg:"id"`
	ExternalID string `pg:"external_id"`

	tableName struct{} `pg:"users"`
}

func (m model) Convert() *entity.User {
	return &entity.User{
		ID:         m.ID,
		ExternalID: m.ExternalID,
	}
}

func newModel(i *entity.User) (*model, error) {
	return &model{
		ID:         i.ID,
		ExternalID: i.ExternalID,
	}, nil
}
