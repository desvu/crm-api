package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID   uint   `pg:"id"`
	Name string `pg:"name"`

	tableName struct{} `pg:"publishers"`
}

func (m model) Convert() *entity.Publisher {
	return &entity.Publisher{
		ID:   m.ID,
		Name: m.Name,
	}
}

func newModel(i *entity.Publisher) (*model, error) {
	return &model{
		ID:   i.ID,
		Name: i.Name,
	}, nil
}
