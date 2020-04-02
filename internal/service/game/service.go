package game

import (
	"context"

	"github.com/pkg/errors"
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
	panic("implement me")
}

func (s Service) Delete(ctx context.Context, id uint) error {
	panic("implement me")
}

func (s Service) Publish(ctx context.Context, id uint) error {
	panic("implement me")
}

func (s Service) GetByID(ctx context.Context, id uint) (*entity.Game, error) {
	panic("implement me")
}

func (s Service) GetExistByID(ctx context.Context, id uint) (*entity.Game, error) {
	game, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if game == nil {
		return nil, errors.WithStack(ErrGameNotFound)
	}

	return game, nil
}
