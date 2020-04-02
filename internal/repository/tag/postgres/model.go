package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID   uint   `pg:"id"`
	Name string `pg:"name"`

	tableName struct{} `pg:"tags"`
}

func (m model) Convert() *entity.Tag {
	return &entity.Tag{
		ID:   m.ID,
		Name: m.Name,
	}
}

func newModel(i *entity.Tag) (*model, error) {
	return &model{
		ID:   i.ID,
		Name: i.Name,
	}, nil
}
