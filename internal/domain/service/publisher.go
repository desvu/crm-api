package service

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/publisher_service.go -package=mocks github.com/qilin/crm-api/internal/domain/service PublisherService
type PublisherService interface {
	Create(ctx context.Context, data *CreatePublisherData) (*entity.Publisher, error)
	Update(ctx context.Context, data *UpdatePublisherData) (*entity.Publisher, error)
	Delete(ctx context.Context, id uint) error

	GetByID(ctx context.Context, id uint) (*entity.Publisher, error)
	GetExistByID(ctx context.Context, id uint) (*entity.Publisher, error)
	GetByIDs(ctx context.Context, ids []uint) ([]entity.Publisher, error)
	GetByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.Publisher, error)

	UpdatePublishersForGameRevision(ctx context.Context, gameRevision *entity.GameRevision, publisherIDs []uint) error
}

type CreatePublisherData struct {
	Name string
}

type UpdatePublisherData struct {
	ID   uint
	Name string
}
