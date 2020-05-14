package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/game_revision_publisher_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository GameRevisionPublisherRepository
type GameRevisionPublisherRepository interface {
	Create(ctx context.Context, i *entity.GameRevisionPublisher) error
	CreateMultiple(ctx context.Context, items []entity.GameRevisionPublisher) error
	Delete(ctx context.Context, i *entity.GameRevisionPublisher) error
	DeleteMultiple(ctx context.Context, items []entity.GameRevisionPublisher) error

	FindByPublisherID(ctx context.Context, publisherID uint) ([]entity.GameRevisionPublisher, error)
	FindByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.GameRevisionPublisher, error)
	FindByGameRevisionIDs(ctx context.Context, gameRevisionIDs []uint) ([]entity.GameRevisionPublisher, error)
	FindByGameRevisionIDAndPublisherIDs(ctx context.Context, gameRevisionID uint, publisherIDs []uint) ([]entity.GameRevisionPublisher, error)
}
