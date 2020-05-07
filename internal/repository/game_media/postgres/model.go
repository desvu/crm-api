package postgres

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/enum/game_media"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID         uint      `pg:"id"`
	GameID     string    `pg:"game_id,notnull,use_zero"`
	Type       uint      `pg:"type,notnull,use_zero"`
	FilePath   string    `pg:"file_path,notnull,use_zero"`
	IsUploaded bool      `pg:"is_uploaded,notnull,use_zero"`
	CreatedAt  time.Time `pg:"created_at"`

	tableName struct{} `pg:"game_media"`
}

func (m model) Convert() *entity.GameMedia {
	return &entity.GameMedia{
		ID:         m.ID,
		GameID:     m.GameID,
		Type:       game_media.NewType(m.Type),
		FilePath:   m.FilePath,
		IsUploaded: m.IsUploaded,
		CreatedAt:  m.CreatedAt,
	}
}

func newModel(i *entity.GameMedia) *model {
	return &model{
		ID:         i.ID,
		GameID:     i.GameID,
		Type:       i.Type.Value(),
		FilePath:   i.FilePath,
		IsUploaded: i.IsUploaded,
		CreatedAt:  i.CreatedAt,
	}
}
