package game_revision

import (
	"context"
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game"
	"github.com/qilin/crm-api/internal/domain/enum/game_revision"
	"github.com/qilin/crm-api/internal/domain/errors"
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

func (s *Service) Update(ctx context.Context, data *service.UpdateGameRevisionData) (*entity.GameRevisionEx, error) {
	revision, err := s.GameRevisionRepository.FindByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if revision == nil {
		return nil, errors.GameRevisionNotFound
	}

	if data.Summary != nil {
		revision.Summary = *data.Summary
	}

	if data.Description != nil {
		revision.Description = *data.Description
	}

	if data.License != nil {
		revision.License = *data.License
	}

	if data.Trailer != nil {
		revision.Trailer = *data.Trailer
	}

	if data.Status != nil {
		if revision.Status == game_revision.StatusDraft || revision.Status == game_revision.StatusPublishing {
			if *data.Status == game_revision.StatusPublished {
				publishedAt := time.Now()
				revision.PublishedAt = &publishedAt
			}
		}
		revision.Status = *data.Status
	}

	if data.ReleaseDate != nil {
		revision.ReleaseDate = *data.ReleaseDate
	}

	if data.PlayTime != nil {
		revision.PlayTime = *data.PlayTime
	}

	if data.Platforms != nil {
		revision.Platforms = *data.Platforms
	}

	if data.SocialLinks != nil {
		var socialLinks []entity.SocialLink
		for _, i := range *data.SocialLinks {
			socialLinks = append(socialLinks, entity.SocialLink{URL: i.URL})
		}
		revision.SocialLinks = socialLinks
	}

	if data.SystemRequirements != nil {
		platforms := map[game.Platform]bool{}
		var systemRequirements []entity.SystemRequirements
		for _, item := range *data.SystemRequirements {
			if platforms[item.Platform] {
				return nil, errors.GameRevisionUniqueSystemRequirements
			}
			set := entity.SystemRequirements{
				Platform: item.Platform,
			}
			if item.Minimal != nil {
				set.Minimal = &entity.RequirementsSet{
					CPU:       item.Minimal.CPU,
					GPU:       item.Minimal.GPU,
					DiskSpace: item.Minimal.DiskSpace,
					RAM:       item.Minimal.RAM,
				}
			}
			if item.Recommended != nil {
				set.Recommended = &entity.RequirementsSet{
					CPU:       item.Recommended.CPU,
					GPU:       item.Recommended.GPU,
					DiskSpace: item.Recommended.DiskSpace,
					RAM:       item.Recommended.RAM,
				}
			}
			systemRequirements = append(systemRequirements, set)
			platforms[item.Platform] = true
		}
		revision.SystemRequirements = systemRequirements
	}

	if err := s.Transactor.Transact(ctx, func(tx context.Context) error {
		if err = s.GameRevisionRepository.Update(tx, revision); err != nil {
			return err
		}

		if data.Tags != nil {
			err := s.TagService.UpdateTagsForGameRevision(tx, revision, *data.Tags)
			if err != nil {
				return err
			}
		}

		if data.Developers != nil {
			err := s.DeveloperService.UpdateDevelopersForGameRevision(tx, revision, *data.Developers)
			if err != nil {
				return err
			}
		}

		if data.Publishers != nil {
			err := s.PublisherService.UpdatePublishersForGameRevision(tx, revision, *data.Publishers)
			if err != nil {
				return err
			}
		}

		if data.Features != nil {
			err := s.FeatureService.UpdateFeaturesForGameRevision(tx, revision, *data.Features)
			if err != nil {
				return err
			}
		}

		if data.Genres != nil {
			err := s.GenreService.UpdateGenresForGameRevision(tx, revision, *data.Genres)
			if err != nil {
				return err
			}
		}

		if data.Media != nil {
			err := s.GameMediaService.UpdateForGameRevision(tx, revision, *data.Media)
			if err != nil {
				return err
			}
		}

		if data.Localizations != nil {
			err := s.LocalizationService.UpdateLocalizationsForGameRevision(tx, revision, *data.Localizations)
			if err != nil {
				return err
			}
		}

		if data.Rating != nil {
			err := s.RatingService.UpdateRatingsForGameRevision(tx, revision, *data.Rating)
			if err != nil {
				return err
			}
		}

		if data.Reviews != nil {
			err := s.ReviewService.UpdateReviewsForGameRevision(tx, revision, *data.Reviews)
			if err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return s.GetByID(ctx, revision.ID)
}

func (s *Service) GetByID(ctx context.Context, id uint) (*entity.GameRevisionEx, error) {
	revision, err := s.GameRevisionExRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if revision == nil {
		return nil, errors.GameRevisionNotFound
	}

	return revision, nil
}

func (s *Service) GetLastByGameIDs(ctx context.Context, gameIDs []string) ([]entity.GameRevisionEx, error) {
	return s.GameRevisionExRepository.FindLastByGameIDs(ctx, gameIDs)
}

func (s *Service) GetByIDAndGameID(ctx context.Context, id uint, gameID string) (*entity.GameRevisionEx, error) {
	revision, err := s.GameRevisionExRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if revision == nil {
		return nil, errors.GameRevisionNotFound
	}

	return revision, nil
}

func (s *Service) GetLastPublishedByGame(ctx context.Context, game *entity.Game) (*entity.GameRevisionEx, error) {
	revision, err := s.GameRevisionRepository.FindLastPublishedByGameID(ctx, game.ID)
	if err != nil {
		return nil, err
	}

	if revision == nil {
		return nil, errors.GameRevisionNotFound
	}

	return s.GameRevisionExRepository.FindByID(ctx, revision.ID)
}

func (s *Service) GetDraftByGame(ctx context.Context, game *entity.Game) (*entity.GameRevisionEx, error) {
	draftRevision, err := s.GameRevisionRepository.FindDraftByGameID(ctx, game.ID)
	if err != nil {
		return nil, err
	}

	if draftRevision != nil {
		return s.GameRevisionExRepository.FindByID(ctx, draftRevision.ID)
	}

	return s.create(ctx, game)
}

func (s *Service) GetByFilter(ctx context.Context, data *service.GetByFilterGameData) ([]entity.GameRevisionEx, error) {
	if data.Limit == 0 {
		data.Limit = 30
	}

	return s.GameRevisionExRepository.FindByFilter(ctx, &repository.FindByFilterGameRevisionData{
		Title:         data.Title,
		OnlyPublished: data.OnlyPublished,
		GenreIDs:      data.GenreIDs,
		FeatureIDs:    data.FeatureIDs,
		Languages:     data.Languages,
		Platforms:     data.Platforms,
		OrderType:     data.OrderType,
		OrderBy:       data.OrderBy,
		Limit:         data.Limit,
		Offset:        data.Offset,
	})
}

func (s *Service) GetCountByFilter(ctx context.Context, data *service.GetByFilterGameData) (int, error) {
	return s.GameRevisionRepository.CountByFilter(ctx, &repository.FindByFilterGameRevisionData{
		Title:         data.Title,
		OnlyPublished: data.OnlyPublished,
		GenreIDs:      data.GenreIDs,
		FeatureIDs:    data.FeatureIDs,
		Languages:     data.Languages,
		Platforms:     data.Platforms,
	})
}

func (s *Service) IsGamesPublished(ctx context.Context, gameIDs []string) error {
	res, err := s.GameRevisionRepository.FindPublishedByGameIDs(ctx, gameIDs)
	if err != nil {
		return err
	}

	if len(res) != len(gameIDs) {
		return errors.GameNotFound // attach game id
	}

	// sort.Strings(res)
	// sort.Strings(ids)

	// for i := range ids {
	// 	if i >= len(res) || res[i] != ids[i] {
	// 		return errors.GameNotFound // attach game id
	// 	}
	// }

	return nil
}
