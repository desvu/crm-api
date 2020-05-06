package game_revision

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game_revision"
	"github.com/qilin/crm-api/internal/domain/errors"
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

	if data.Status != nil {
		revision.Status = *data.Status
	}

	if data.ReleaseDate != nil {
		revision.ReleaseDate = *data.ReleaseDate
	}

	if data.Platforms != nil {
		revision.Platforms = *data.Platforms
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

	newRevision := &entity.GameRevision{
		GameID: game.ID,
		Status: game_revision.StatusDraft,
	}

	if err = s.GameRevisionRepository.Create(ctx, newRevision); err != nil {
		return nil, err
	}

	return &entity.GameRevisionEx{
		GameRevision: *newRevision,
	}, nil
}

func (s *Service) IsGamesPublished(ctx context.Context, ids ...string) error {
	if len(ids) == 0 {
		return nil
	}
	res, err := s.GameRevisionRepository.GetPublishedIds(ctx, ids...)
	if err != nil {
		return err
	}

	if len(res) != len(ids) {
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
