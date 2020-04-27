package game_store_publish

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game_publish"
	"github.com/qilin/crm-api/internal/domain/errors"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

func (s Service) Create(ctx context.Context, data *service.CreateGameStorePublishData) (*entity.GameStorePublish, error) {
	gsp := &entity.GameStorePublish{
		GameID: data.GameID,
		Body:   data.Body,
		Status: game_publish.StatusInProcess,
	}

	if err := s.GameStorePublishRepository.Create(ctx, gsp); err != nil {
		return nil, err
	}

	return gsp, nil
}

func (s Service) Update(ctx context.Context, data *service.UpdateGameStorePublishData) (*entity.GameStorePublish, error) {
	gsp, err := s.GetExistByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	gsp.Status = data.Status
	if err := s.GameStorePublishRepository.Update(ctx, gsp); err != nil {
		return nil, err
	}

	return gsp, nil
}

func (s Service) GetByID(ctx context.Context, id uint) (*entity.GameStorePublish, error) {
	return s.GameStorePublishRepository.FindByID(ctx, id)
}

func (s Service) GetExistByID(ctx context.Context, id uint) (*entity.GameStorePublish, error) {
	gsp, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if gsp == nil {
		return nil, errors.GameStorePublishNotFound
	}

	return gsp, nil
}
