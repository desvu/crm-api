package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type PublisherService interface {
	Create(ctx context.Context, data *CreatePublisherData) (*entity.Publisher, error)
	Update(ctx context.Context, data *UpdatePublisherData) (*entity.Publisher, error)
	Delete(ctx context.Context, id uint) error

	GetByID(ctx context.Context, id uint) (*entity.Publisher, error)
	GetExistByID(ctx context.Context, id uint) (*entity.Publisher, error)
	GetByIDs(ctx context.Context, ids []uint) ([]entity.Publisher, error)
	GetByGameID(ctx context.Context, gameID uint) ([]entity.Publisher, error)

	UpdatePublishersForGame(ctx context.Context, game *entity.Game, publisherIDs []uint) error
}

type CreatePublisherData struct {
	Name string
}

type UpdatePublisherData struct {
	ID   uint
	Name string
}
