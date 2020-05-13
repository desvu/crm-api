package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game_rating"
)

type model struct {
	ID                  uint  `pg:"id"`
	GameRevisionID      uint  `pg:"game_revision_id"`
	Agency              uint8 `pg:"agency"`
	Rating              uint8 `pg:"rating"`
	DisplayOnlineNotice bool  `pg:"display_online_notice"`
	ShowAgeRestrict     bool  `pg:"show_age_restrict"`

	tableName struct{} `pg:"game_revision_ratings"`
}

func (m model) Convert() *entity.Rating {
	return &entity.Rating{
		ID:                  m.ID,
		GameRevisionID:      m.GameRevisionID,
		Agency:              game_rating.NewAgency(m.Agency),
		Rating:              game_rating.NewRating(m.Rating),
		DisplayOnlineNotice: m.DisplayOnlineNotice,
		ShowAgeRestrict:     m.ShowAgeRestrict,
	}
}

func newModel(i *entity.Rating) (*model, error) {
	return &model{
		ID:                  i.ID,
		GameRevisionID:      i.GameRevisionID,
		Agency:              i.Agency.Value(),
		Rating:              i.Rating.Value(),
		DisplayOnlineNotice: i.DisplayOnlineNotice,
		ShowAgeRestrict:     i.ShowAgeRestrict,
	}, nil
}
