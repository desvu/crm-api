package game_revision

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/service"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game_revision"
	"github.com/qilin/crm-api/internal/domain/errors"
)

func (s *Service) create(ctx context.Context, game *entity.Game) (*entity.GameRevisionEx, error) {
	lastDraftRevision, err := s.GameRevisionRepository.FindDraftByGameID(ctx, game.ID)
	if err != nil {
		return nil, err
	}

	if lastDraftRevision != nil {
		return nil, errors.GameRevisionDraftAlreadyExist
	}

	lastPublishedRevision, err := s.GameRevisionExRepository.FindLastPublishedByGameID(ctx, game.ID)
	if err != nil {
		return nil, err
	}

	if lastPublishedRevision == nil {
		return s.createEmptyDraft(ctx, game)
	}

	return s.createDraftFromPublished(ctx, lastPublishedRevision)
}

func (s *Service) createEmptyDraft(ctx context.Context, game *entity.Game) (*entity.GameRevisionEx, error) {
	newRevision := &entity.GameRevision{
		GameID:             game.ID,
		Status:             game_revision.StatusDraft,
		SystemRequirements: []entity.SystemRequirements{},
	}

	if err := s.GameRevisionRepository.Create(ctx, newRevision); err != nil {
		return nil, err
	}

	return &entity.GameRevisionEx{
		GameRevision: *newRevision,
	}, nil
}

func (s *Service) createDraftFromPublished(ctx context.Context, publishedRevision *entity.GameRevisionEx) (*entity.GameRevisionEx, error) {
	newRevision := &entity.GameRevision{
		GameID:             publishedRevision.GameID,
		Summary:            publishedRevision.Summary,
		Description:        publishedRevision.Description,
		License:            publishedRevision.License,
		Trailer:            publishedRevision.Trailer,
		PlayTime:           publishedRevision.PlayTime,
		ReleaseDate:        publishedRevision.ReleaseDate,
		SystemRequirements: publishedRevision.SystemRequirements,
		SocialLinks:        publishedRevision.SocialLinks,
		Platforms:          publishedRevision.Platforms,
		Status:             game_revision.StatusDraft,
	}

	if err := s.Transactor.Transact(ctx, func(tx context.Context) error {
		if err := s.GameRevisionRepository.Create(tx, newRevision); err != nil {
			return err
		}

		err := s.FeatureService.UpdateFeaturesForGameRevision(
			tx,
			newRevision,
			entity.NewFeatureArray(publishedRevision.Features).IDs(),
		)

		if err != nil {
			return err
		}

		err = s.GenreService.UpdateGenresForGameRevision(
			tx,
			newRevision,
			entity.NewGenreArray(publishedRevision.Genres).IDs(),
		)

		if err != nil {
			return err
		}

		err = s.DeveloperService.UpdateDevelopersForGameRevision(
			tx,
			newRevision,
			entity.NewDeveloperArray(publishedRevision.Developers).IDs(),
		)

		if err != nil {
			return err
		}

		err = s.PublisherService.UpdatePublishersForGameRevision(
			tx,
			newRevision,
			entity.NewPublisherArray(publishedRevision.Publishers).IDs(),
		)

		if err != nil {
			return err
		}

		err = s.TagService.UpdateTagsForGameRevision(
			tx,
			newRevision,
			entity.NewTagArray(publishedRevision.Tags).IDs(),
		)

		if err != nil {
			return err
		}

		err = s.GameMediaService.UpdateForGameRevision(
			tx,
			newRevision,
			entity.NewGameMediaArray(publishedRevision.Media).IDs(),
		)

		if err != nil {
			return err
		}

		var localizationData []service.LocalizationData
		for _, localization := range publishedRevision.Localization {
			localizationData = append(localizationData, service.LocalizationData{
				Language:  localization.Language,
				Interface: localization.Interface,
				Audio:     localization.Audio,
				Subtitles: localization.Subtitles,
			})
		}

		err = s.LocalizationService.UpdateLocalizationsForGameRevision(tx, newRevision, localizationData)
		if err != nil {
			return err
		}

		var ratingData []service.RatingData
		for _, rating := range publishedRevision.Rating {
			ratingData = append(ratingData, service.RatingData{
				Agency:              rating.Agency,
				Rating:              rating.Rating,
				DisplayOnlineNotice: rating.DisplayOnlineNotice,
				ShowAgeRestrict:     rating.ShowAgeRestrict,
			})
		}

		err = s.RatingService.UpdateRatingsForGameRevision(tx, newRevision, ratingData)
		if err != nil {
			return err
		}

		var reviewData []service.ReviewData
		for _, rating := range publishedRevision.Rating {
			ratingData = append(ratingData, service.RatingData{
				Agency:              rating.Agency,
				Rating:              rating.Rating,
				DisplayOnlineNotice: rating.DisplayOnlineNotice,
				ShowAgeRestrict:     rating.ShowAgeRestrict,
			})
		}

		err = s.ReviewService.UpdateReviewsForGameRevision(tx, newRevision, reviewData)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &entity.GameRevisionEx{
		GameRevision: *newRevision,
		Tags:         publishedRevision.Tags,
		Developers:   publishedRevision.Developers,
		Publishers:   publishedRevision.Publishers,
		Features:     publishedRevision.Features,
		Genres:       publishedRevision.Genres,
		Media:        publishedRevision.Media,
		Localization: publishedRevision.Localization,
		Rating:       publishedRevision.Rating,
		Review:       publishedRevision.Review,
	}, nil
}
