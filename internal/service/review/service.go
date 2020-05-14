package review

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

func (s Service) GetByGameRevisionID(ctx context.Context, gameID uint) ([]entity.Review, error) {
	return s.ReviewRepository.FindByGameRevisionID(ctx, gameID)
}

func (s Service) UpdateReviewsForGameRevision(ctx context.Context, gameRevision *entity.GameRevision, reviews []service.ReviewData) error {
	currentGameReviews, err := s.ReviewRepository.FindByGameRevisionID(ctx, gameRevision.ID)
	if err != nil {
		return err
	}

	var currentReviews []entity.Review
	for _, review := range currentGameReviews {
		for _, newReview := range reviews {
			if newReview.Link == review.Link {
				currentReviews = append(currentReviews, review)
			}
		}
	}

	return s.Transactor.Transact(ctx, func(tx context.Context) error {
		err = s.ReviewRepository.DeleteMultiple(ctx, getGameReviewForDelete(currentReviews, currentGameReviews))
		if err != nil {
			return err
		}

		err = s.ReviewRepository.CreateMultiple(ctx, getGameReviewForInsert(gameRevision.ID, reviews, currentGameReviews))
		if err != nil {
			return err
		}

		err = s.ReviewRepository.UpdateMultiple(ctx, getGameReviewForUpdate(gameRevision.ID, reviews, currentGameReviews))
		if err != nil {
			return err
		}

		return nil
	})
}

func getGameReviewForInsert(gameID uint, newReviews []service.ReviewData, currentGameReviews []entity.Review) []entity.Review {
	gameReviews := make([]entity.Review, 0)
	for _, newReview := range newReviews {
		var hasMatch bool
		for _, currentGameReview := range currentGameReviews {
			if newReview.Link == currentGameReview.Link {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameReviews = append(gameReviews, entity.Review{
				GameRevisionID: gameID,
				PressName:      newReview.PressName,
				Link:           newReview.Link,
				Score:          newReview.Score,
				Quote:          newReview.Quote,
			})
		}
	}
	return gameReviews
}

func getGameReviewForUpdate(gameID uint, newReviews []service.ReviewData, currentGameReviews []entity.Review) []entity.Review {
	gameReviews := make([]entity.Review, 0)
	for _, newReview := range newReviews {
		for _, currentGameReview := range currentGameReviews {
			if newReview.Link == currentGameReview.Link {
				continue
			}

			if newReview.Quote != currentGameReview.Quote || newReview.Score != currentGameReview.Score ||
				newReview.PressName != currentGameReview.PressName {
				gameReviews = append(gameReviews, entity.Review{
					ID:             currentGameReview.ID,
					GameRevisionID: gameID,
					PressName:      currentGameReview.PressName,
					Link:           currentGameReview.Link,
					Score:          currentGameReview.Score,
					Quote:          currentGameReview.Quote,
				})
			}
		}
	}
	return gameReviews
}

func getGameReviewForDelete(reviews []entity.Review, currentGameReviews []entity.Review) []entity.Review {
	gameReviews := make([]entity.Review, 0)
	for _, currentGameReview := range currentGameReviews {
		var hasMatch bool
		for _, newReview := range reviews {
			if currentGameReview.Link == newReview.Link {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameReviews = append(gameReviews, entity.Review{
				ID:             currentGameReview.ID,
				GameRevisionID: currentGameReview.GameRevisionID,
				PressName:      currentGameReview.PressName,
				Link:           currentGameReview.Link,
				Score:          currentGameReview.Score,
				Quote:          currentGameReview.Quote,
			})
		}
	}
	return gameReviews
}
