package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/enum/game_publish"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type GameStorePublishService interface {
	Create(ctx context.Context, data *CreateGameData) (*entity.GameStorePublish, error)
	Update(ctx context.Context, data *UpdateGameData) (*entity.GameStorePublish, error)
}

type CreateGameStorePublishData struct {
	GameID string
	Body   string
}

type UpdateGameStorePublishData struct {
	ID     uint
	Status game_publish.Status
}
