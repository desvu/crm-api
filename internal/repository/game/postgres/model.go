package postgres

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game"
)

type model struct {
	ID          uint      `pg:"id"`
	Title       string    `pg:"title,notnull,use_zero"`
	Summary     string    `pg:"summary,notnull,use_zero"`
	Description string    `pg:"description,notnull,use_zero"`
	License     string    `pg:"license,notnull,use_zero"`
	Ranking     string    `pg:"ranking,notnull,use_zero"`
	Type        uint8     `pg:"type,notnull,use_zero"`
	Platforms   []uint8   `pg:"platforms,array,notnull,use_zero"`
	ReleaseDate time.Time `pg:"release_date,notnull,use_zero"`

	tableName struct{} `pg:"games"`
}

func (m model) Convert() *entity.Game {
	return &entity.Game{
		ID:          m.ID,
		Title:       m.Title,
		Summary:     m.Summary,
		Description: m.Description,
		License:     m.License,
		Ranking:     m.Ranking,
		Type:        game.NewType(m.Type),
		Platforms:   game.NewPlatformArray(m.Platforms...),
		ReleaseDate: m.ReleaseDate,
	}
}

func newModel(i *entity.Game) (*model, error) {
	return &model{
		ID:          i.ID,
		Title:       i.Title,
		Summary:     i.Summary,
		Description: i.Description,
		License:     i.License,
		Ranking:     i.Ranking,
		Type:        i.Type.Value(),
		Platforms:   i.Platforms.Values(),
		ReleaseDate: i.ReleaseDate,
	}, nil
}
