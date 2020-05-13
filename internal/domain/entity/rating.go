package entity

import "github.com/qilin/crm-api/internal/domain/enum/game_rating"

type Rating struct {
	ID                  uint
	GameRevisionID      uint
	Agency              game_rating.Agency
	Rating              game_rating.Rating
	DisplayOnlineNotice bool
	ShowAgeRestrict     bool
}
