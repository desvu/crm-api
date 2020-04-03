package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID   uint   `pg:"id"`
	Name string `pg:"name"`

	tableName struct{} `pg:"features"`
}

func (m model) Convert() *entity.Feature {
	return &entity.Feature{
		ID:   m.ID,
		Name: m.Name,
	}
}

func newModel(i *entity.Feature) (*model, error) {
	return &model{
		ID:   i.ID,
		Name: i.Name,
	}, nil
}
