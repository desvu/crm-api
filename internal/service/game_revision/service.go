package game_revision

import (
	"context"
	"errors"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

var ErrGameRevisionNotFound = errors.New("game revision not found")

func (s Service) Update(ctx context.Context, data *service.UpdateGameRevisionData) (*entity.GameRevisionEx, error) {
	revision, err := s.GameRevisionRepository.FindByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if revision == nil {
		return nil, ErrGameRevisionNotFound
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

func (s Service) GetByID(ctx context.Context, id uint) (*entity.GameRevisionEx, error) {
	return s.GameRevisionExRepository.FindByID(ctx, id)
}

func (s Service) GetExistByID(ctx context.Context, id uint) (*entity.GameRevisionEx, error) {
	revision, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if revision == nil {
		return nil, ErrGameRevisionNotFound
	}

	return revision, nil
}

func (s Service) GetDraftByGame(ctx context.Context, game *entity.Game) (*entity.GameRevisionEx, error) {
	draftRevision, err := s.GameRevisionRepository.FindDraftByGameID(ctx, game.ID)
	if err != nil {
		return nil, err
	}

	if draftRevision != nil {
		return s.GameRevisionExRepository.FindByID(ctx, draftRevision.ID)
	}

	newRevision := &entity.GameRevision{
		GameID: game.ID,
	}

	if err = s.GameRevisionRepository.Create(ctx, newRevision); err != nil {
		return nil, err
	}

	return &entity.GameRevisionEx{
		GameRevision: *newRevision,
	}, nil
}
