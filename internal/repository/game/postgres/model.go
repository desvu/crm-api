package postgres

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game"
)

type model struct {
	ID          uint      `pg:"id"`
	Title       string    `pg:"title"`
	Summary     string    `pg:"summary"`
	Description string    `pg:"description"`
	License     string    `pg:"license"`
	Ranking     string    `pg:"ranking"`
	Type        uint8     `pg:"type"`
	Platforms   []uint8   `pg:"platforms,array"`
	ReleaseDate time.Time `pg:"release_date"`

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
		Platforms:   game.NewPlatformArray(m.Platforms),
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
