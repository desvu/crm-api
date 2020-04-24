package game

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game_revision"
	"github.com/qilin/crm-api/internal/domain/publisher"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

var ErrGameNotFound = errors.New("game not found")

func (s Service) Create(ctx context.Context, data *service.CreateGameData) (*entity.GameEx, error) {
	game := &entity.Game{
		ID:    uuid.New().String(),
		Title: data.Title,
		Type:  data.Type,
	}

	var updatedRevision *entity.GameRevisionEx
	if err := s.Transactor.Transact(ctx, func(tx context.Context) error {
		if err := s.GameRepository.Create(tx, game); err != nil {
			return err
		}

		revision, err := s.GameRevisionService.GetDraftByGame(tx, game)
		if err != nil {
			return err
		}

		updatedRevision, err = s.GameRevisionService.Update(tx, &service.UpdateGameRevisionData{
			ID:          revision.ID,
			Summary:     data.Summary,
			Description: data.Description,
			Slug:        data.Slug,
			License:     data.License,
			Tags:        data.Tags,
			Developers:  data.Developers,
			Publishers:  data.Publishers,
			Features:    data.Features,
			Genres:      data.Genres,
			ReleaseDate: data.ReleaseDate,
			Platforms:   data.Platforms,
		})

		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &entity.GameEx{
		Game:     *game,
		Revision: updatedRevision,
	}, nil
}

func (s Service) Update(ctx context.Context, data *service.UpdateGameData) (*entity.GameEx, error) {
	game, err := s.GetByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	revision, err := s.GameRevisionService.GetDraftByGame(ctx, game)
	if err != nil {
		return nil, err
	}

	var updatedRevision *entity.GameRevisionEx
	if err := s.Transactor.Transact(ctx, func(ctx context.Context) error {
		if err = s.GameRepository.Update(ctx, game); err != nil {
			return err
		}

		updatedRevision, err = s.GameRevisionService.Update(ctx, &service.UpdateGameRevisionData{
			ID:         revision.ID,
			Tags:       data.Tags,
			Developers: data.Developers,
			Publishers: data.Publishers,
			Features:   data.Features,
			Genres:     data.Genres,
		})

		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &entity.GameEx{
		Game:     *game,
		Revision: updatedRevision,
	}, nil
}

func (s Service) Delete(ctx context.Context, id string) error {
	game, err := s.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return s.GameRepository.Delete(ctx, game)
}

func (s Service) Publish(ctx context.Context, id string) error {
	game, err := s.GetExByID(ctx, id)
	if err != nil {
		return err
	}

	if err = s.GameStorePublisher.Publish(publisher.PublishGameStoreData{Game: game}); err != nil {
		return err
	}

	revisionStatus := game_revision.StatusPublishing
	_, err = s.GameRevisionService.Update(ctx, &service.UpdateGameRevisionData{
		ID:     game.Revision.ID,
		Status: &revisionStatus,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s Service) GetByID(ctx context.Context, id string) (*entity.Game, error) {
	game, err := s.GameRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, ErrGameNotFound
	}

	return game, nil
}

func (s Service) GetExByID(ctx context.Context, id string) (*entity.GameEx, error) {
	game, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	revision, err := s.GameRevisionService.GetDraftByGame(ctx, game)
	if err != nil {
		return nil, err
	}

	return &entity.GameEx{
		Game:     *game,
		Revision: revision,
	}, nil
}
