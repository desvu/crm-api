package postgres

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game"
	"github.com/qilin/crm-api/internal/domain/enum/game_revision"
)

type model struct {
	ID          uint       `pg:"id"`
	GameID      string     `pg:"game_id,notnull,use_zero"`
	Summary     string     `pg:"summary,notnull,use_zero"`
	Description string     `pg:"description,notnull,use_zero"`
	License     string     `pg:"license,notnull,use_zero"`
	Status      uint8      `pg:"status,notnull,use_zero"`
	Platforms   []uint8    `pg:"platforms,array,notnull,use_zero"`
	ReleaseDate time.Time  `pg:"release_date,notnull,use_zero"`
	PublishedAt *time.Time `pg:"published_at"`

	tableName struct{} `pg:"game_revisions"`
}

func (m model) Convert() *entity.GameRevision {
	return &entity.GameRevision{
		ID:          m.ID,
		GameID:      m.GameID,
		Summary:     m.Summary,
		Description: m.Description,
		License:     m.License,
		Status:      game_revision.NewStatus(m.Status),
		Platforms:   game.NewPlatformArray(m.Platforms...),
		ReleaseDate: m.ReleaseDate,
		PublishedAt: m.PublishedAt,
	}
}

func newModel(i *entity.GameRevision) (*model, error) {
	return &model{
		ID:          i.ID,
		GameID:      i.GameID,
		Summary:     i.Summary,
		Description: i.Description,
		License:     i.License,
		Status:      i.Status.Value(),
		Platforms:   i.Platforms.Values(),
		ReleaseDate: i.ReleaseDate,
		PublishedAt: i.PublishedAt,
	}, nil
}
