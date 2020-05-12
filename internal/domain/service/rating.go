package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game_rating"
	"github.com/qilin/crm-api/internal/domain/errors"
)

//go:generate mockgen -destination=../mocks/rating_service.go -package=mocks github.com/qilin/crm-api/internal/domain/service RatingService
type RatingService interface {
	GetByID(ctx context.Context, id uint) (*entity.Rating, error)
	GetExistByID(ctx context.Context, id uint) (*entity.Rating, error)
	GetByIDs(ctx context.Context, ids []uint) ([]entity.Rating, error)
	GetByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.Rating, error)

	UpdateRatingsForGameRevision(ctx context.Context, gameRevision *entity.GameRevision, ratings []RatingData) error
}

type RatingData struct {
	Agency              game_rating.Agency
	Rating              game_rating.Rating
	DisplayOnlineNotice bool
	ShowAgeRestrict     bool
}

func (d RatingData) Validate() error {
	if d.Agency.Value() == 0 {
		return errors.RatingUndefinedAgency
	}
	if d.Rating.Value() == 0 {
		return errors.RatingUndefinedRating
	}
	return nil
}
