package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game"
)

type model struct {
	ID   uint   `pg:"id"`
	Name string `pg:"name"`
	Icon string `pg:"icon"`

	tableName struct{} `pg:"features"`
}

func (m model) Convert() *entity.Feature {
	return &entity.Feature{
		ID:   m.ID,
		Name: m.Name,
		Icon: game.Icon(m.Icon),
	}
}

func newModel(i *entity.Feature) (*model, error) {
	return &model{
		ID:   i.ID,
		Name: i.Name,
		Icon: i.Icon.String(),
	}, nil
}
