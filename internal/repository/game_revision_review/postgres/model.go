package postgres

import (
	"github.com/qilin/crm-api/internal/domain/entity"
)

type model struct {
	ID             uint   `pg:"id"`
	GameRevisionID uint   `pg:"game_revision_id"`
	PressName      string `pg:"press_name"`
	Link           string `pg:"link"`
	Score          uint   `pg:"score"`
	Quote          string `pg:"quote"`

	tableName struct{} `pg:"game_revision_reviews"`
}

func (m model) Convert() *entity.Review {
	return &entity.Review{
		ID:             m.ID,
		GameRevisionID: m.GameRevisionID,
		PressName:      m.PressName,
		Link:           m.Link,
		Score:          m.Score,
		Quote:          m.Quote,
	}
}

func newModel(i *entity.Review) (*model, error) {
	return &model{
		ID:             i.ID,
		GameRevisionID: i.GameRevisionID,
		PressName:      i.PressName,
		Link:           i.Link,
		Score:          i.Score,
		Quote:          i.Quote,
	}, nil
}
