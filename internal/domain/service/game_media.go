package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/game_media"
	"github.com/qilin/crm-api/pkg/errors"
)

var ErrGameMediaNotFound = errors.NewService(errors.ErrNotFound, "game media not found")

type GameMediaService interface {
	Create(ctx context.Context, data *CreateGameMediaData) (*entity.GameMedia, error)
	Upload(ctx context.Context, data *UploadGameMediaData) (*entity.GameMedia, error)
	Delete(ctx context.Context, id uint) error

	GetByID(ctx context.Context, id uint) (*entity.GameMedia, error)
	GetByIDs(ctx context.Context, ids []uint) ([]entity.GameMedia, error)
	GetByRevision(ctx context.Context, revision *entity.GameRevision) ([]entity.GameMedia, error)
}

type CreateGameMediaData struct {
	GameID    string
	Type      game_media.Type
	Extension string
}

type UploadGameMediaData struct {
	ID     uint
	GameID string
	Image  []byte
}

type UpdateGameMediaData struct {
	ID           uint
	GameRevision *entity.GameRevision
}
