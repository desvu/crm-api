package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID             uint   `pg:"id"`
	GameRevisionID uint   `pg:"game_revision_id"`
	Language       string `pg:"language"`
	Interface      bool   `pg:"interface"`
	Audio          bool   `pg:"audio"`
	Subtitles      bool   `pg:"subtitles"`

	tableName struct{} `pg:"game_revision_localizations"`
}

func (m model) Convert() *entity.Localization {
	return &entity.Localization{
		ID:             m.ID,
		GameRevisionID: m.GameRevisionID,
		Language:       m.Language,
		Interface:      m.Interface,
		Audio:          m.Audio,
		Subtitles:      m.Subtitles,
	}
}

func newModel(i *entity.Localization) (*model, error) {
	return &model{
		ID:             i.ID,
		GameRevisionID: i.GameRevisionID,
		Language:       i.Language,
		Interface:      i.Interface,
		Audio:          i.Audio,
		Subtitles:      i.Subtitles,
	}, nil
}
