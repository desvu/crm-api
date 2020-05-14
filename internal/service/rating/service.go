package rating

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/errors"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

func (s *Service) GetByID(ctx context.Context, id uint) (*entity.Rating, error) {
	return s.RatingRepository.FindByID(ctx, id)
}

func (s *Service) GetExistByID(ctx context.Context, id uint) (*entity.Rating, error) {
	rating, err := s.RatingRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if rating == nil {
		return nil, errors.RatingNotFound
	}

	return rating, nil
}

func (s *Service) GetByIDs(ctx context.Context, ids []uint) ([]entity.Rating, error) {
	return s.RatingRepository.FindByIDs(ctx, ids)
}

func (s *Service) GetByGameRevisionID(ctx context.Context, gameID uint) ([]entity.Rating, error) {
	return s.RatingRepository.FindByGameRevisionID(ctx, gameID)
}

func (s *Service) UpdateRatingsForGameRevision(ctx context.Context, gameRevision *entity.GameRevision, ratings []service.RatingData) error {
	agencies := make([]uint8, len(ratings))
	for i, l := range ratings {
		agencies[i] = l.Agency.Value()
	}

	currentRatings, err := s.RatingRepository.FindByGameRevisionIDAndAgency(ctx, gameRevision.ID, agencies)
	if err != nil {
		return err
	}

	currentGameRatings, err := s.RatingRepository.FindByGameRevisionID(ctx, gameRevision.ID)
	if err != nil {
		return err
	}

	err = s.RatingRepository.DeleteMultiple(ctx, getGameRatingsForDelete(currentRatings, currentGameRatings))
	if err != nil {
		return err
	}

	err = s.RatingRepository.CreateMultiple(ctx, getGameRatingsForInsert(gameRevision.ID, ratings, currentGameRatings))
	if err != nil {
		return err
	}

	err = s.RatingRepository.UpdateMultiple(ctx, getGameRatingsForUpdate(gameRevision.ID, ratings, currentGameRatings))
	if err != nil {
		return err
	}

	return nil
}

func getGameRatingsForInsert(gameID uint, newRatings []service.RatingData, currentGameRatings []entity.Rating) []entity.Rating {
	gameRatings := make([]entity.Rating, 0)
	for _, newRating := range newRatings {
		var hasMatch bool
		for _, currentGameRating := range currentGameRatings {
			if newRating.Agency.Value() == currentGameRating.Agency.Value() && currentGameRating.Rating.Value() == newRating.Rating.Value() {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameRatings = append(gameRatings, entity.Rating{
				GameRevisionID:      gameID,
				Agency:              newRating.Agency,
				Rating:              newRating.Rating,
				DisplayOnlineNotice: newRating.DisplayOnlineNotice,
				ShowAgeRestrict:     newRating.ShowAgeRestrict,
			})
		}
	}
	return gameRatings
}

func getGameRatingsForUpdate(gameID uint, newRatings []service.RatingData, currentGameRatings []entity.Rating) []entity.Rating {
	gameRatings := make([]entity.Rating, 0)
	for _, newRating := range newRatings {
		for _, currentGameRating := range currentGameRatings {
			if newRating.Agency.Value() != currentGameRating.Agency.Value() && currentGameRating.Rating.Value() != newRating.Rating.Value() {
				continue
			}

			if newRating.Rating != currentGameRating.Rating || newRating.DisplayOnlineNotice != currentGameRating.DisplayOnlineNotice ||
				newRating.ShowAgeRestrict != currentGameRating.ShowAgeRestrict {
				gameRatings = append(gameRatings, entity.Rating{
					ID:                  currentGameRating.ID,
					GameRevisionID:      gameID,
					Agency:              currentGameRating.Agency,
					Rating:              newRating.Rating,
					DisplayOnlineNotice: newRating.DisplayOnlineNotice,
					ShowAgeRestrict:     newRating.ShowAgeRestrict,
				})
			}
		}
	}
	return gameRatings
}

func getGameRatingsForDelete(ratings []entity.Rating, currentGameRatings []entity.Rating) []entity.Rating {
	gameRatings := make([]entity.Rating, 0)
	for _, currentGameRating := range currentGameRatings {
		var hasMatch bool
		for _, newRating := range ratings {
			if currentGameRating.Agency.Value() == newRating.Agency.Value() && currentGameRating.Rating.Value() == newRating.Rating.Value() {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameRatings = append(gameRatings, entity.Rating{
				ID:                  currentGameRating.ID,
				GameRevisionID:      currentGameRating.GameRevisionID,
				Agency:              currentGameRating.Agency,
				Rating:              currentGameRating.Rating,
				DisplayOnlineNotice: currentGameRating.DisplayOnlineNotice,
				ShowAgeRestrict:     currentGameRating.ShowAgeRestrict,
			})
		}
	}
	return gameRatings
}
