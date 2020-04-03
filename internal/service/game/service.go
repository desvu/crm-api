package game

import (
	"context"
	"errors"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

var ErrGameNotFound = errors.New("game not found")

func (s Service) Create(ctx context.Context, data *service.CreateGameData) (*entity.Game, error) {
	game := &entity.Game{
		Title:       data.Title,
		Summary:     data.Summary,
		Description: data.Description,
		License:     data.License,
		Ranking:     data.Ranking,
		Type:        data.Type,
		Platforms:   data.Platforms,
		ReleaseDate: data.ReleaseDate,
	}

	if err := s.GameRepository.Create(ctx, game); err != nil {
		return nil, err
	}

	return game, nil
}

func (s Service) Update(ctx context.Context, data *service.UpdateGameData) (*entity.Game, error) {
	game, err := s.GetExistByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	game.Title = data.Title
	if err = s.GameRepository.Update(ctx, game); err != nil {
		return nil, err
	}

	return game, nil
}

func (s Service) UpdateEx(ctx context.Context, data *service.UpdateGameExData) (*entity.GameEx, error) {
	game, err := s.Update(ctx, &data.UpdateGameData)
	if err != nil {
		return nil, err
	}

	if data.Tags != nil {
		err := s.TagService.UpdateTagsForGame(ctx, game, *data.Tags)
		if err != nil {
			return nil, err
		}
	}

	if data.Developers != nil {
		err := s.DeveloperService.UpdateDevelopersForGame(ctx, game, *data.Developers)
		if err != nil {
			return nil, err
		}
	}

	if data.Publishers != nil {
		err := s.PublisherService.UpdatePublishersForGame(ctx, game, *data.Publishers)
		if err != nil {
			return nil, err
		}
	}

	if data.Features != nil {
		err := s.FeatureService.UpdateFeaturesForGame(ctx, game, *data.Features)
		if err != nil {
			return nil, err
		}
	}

	if data.Genres != nil {
		err := s.GenreService.UpdateGenreForGame(ctx, game, *data.Genres)
		if err != nil {
			return nil, err
		}
	}

	return s.GameExRepository.FindByID(ctx, game.ID)
}

func (s Service) Delete(ctx context.Context, id uint) error {
	game, err := s.GetExistByID(ctx, id)
	if err != nil {
		return err
	}

	return s.GameRepository.Delete(ctx, game)
}

func (s Service) Publish(ctx context.Context, id uint) error {
	panic("implement me") // TODO
}

func (s Service) GetByID(ctx context.Context, id uint) (*entity.Game, error) {
	return s.GameRepository.FindByID(ctx, id)
}

func (s Service) GetExistByID(ctx context.Context, id uint) (*entity.Game, error) {
	game, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, ErrGameNotFound
	}

	return game, nil
}
