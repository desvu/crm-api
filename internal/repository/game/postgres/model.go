package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game"
)

type model struct {
	ID    string `pg:"id"`
	Title string `pg:"title,notnull,use_zero"`
	Type  uint8  `pg:"type,notnull,use_zero"`

	tableName struct{} `pg:"games"`
}

func (m model) Convert() *entity.Game {
	return &entity.Game{
		ID:    m.ID,
		Title: m.Title,
		Type:  game.NewType(m.Type),
	}
}

func newModel(i *entity.Game) (*model, error) {
	return &model{
		ID:    i.ID,
		Title: i.Title,
		Type:  i.Type.Value(),
	}, nil
}
