package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID   uint   `pg:"id"`
	Name string `pg:"name"`

	tableName struct{} `pg:"developers"`
}

func (m model) Convert() *entity.Developer {
	return &entity.Developer{
		ID:   m.ID,
		Name: m.Name,
	}
}

func newModel(i *entity.Developer) (*model, error) {
	return &model{
		ID:   i.ID,
		Name: i.Name,
	}, nil
}
