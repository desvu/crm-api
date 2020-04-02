package game

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/service"
)

func (s Service) Create(ctx context.Context, data *service.CreateGameData) (*entity.Game, error) {
	gm := &entity.Game{
		Title:       data.Title,
		Summary:     data.Summary,
		Description: data.Description,
		License:     data.License,
		Ranking:     data.Ranking,
		Type:        data.Type,
		Platforms:   data.Platforms,
		ReleaseDate: data.ReleaseDate,
	}

	if err := s.GameRepository.Create(ctx, gm); err != nil {
		return nil, err
	}

	return gm, nil
}

func (s Service) Update(ctx context.Context, data *service.UpdateGameData) (*entity.Game, error) {
	gm, err := s.GameRepository.FindByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}
	if data.Title != nil {
		gm.Title = *data.Title
	}
	err = s.GameRepository.Update(ctx, gm)
	return gm, err
}

func (s Service) Delete(ctx context.Context, id uint) error {
	return s.GameRepository.Delete(ctx, id)
}

func (s Service) Publish(ctx context.Context, id uint) error {
	panic("implement me") // TODO
}

func (s Service) GetByID(ctx context.Context, id uint) (*entity.Game, error) {
	return s.GameRepository.FindByID(ctx, id)
}
