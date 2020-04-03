package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID   uint   `pg:"id"`
	Name string `pg:"name"`

	tableName struct{} `pg:"genres"`
}

func (m model) Convert() *entity.Genre {
	return &entity.Genre{
		ID:   m.ID,
		Name: m.Name,
	}
}

func newModel(i *entity.Genre) (*model, error) {
	return &model{
		ID:   i.ID,
		Name: i.Name,
	}, nil
}
