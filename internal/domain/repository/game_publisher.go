package repository

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
)

//go:generate mockgen -destination=../mocks/game_publisher_repository.go -package=mocks github.com/qilin/crm-api/internal/domain/repository GamePublisherRepository
type GamePublisherRepository interface {
	Create(ctx context.Context, i *entity.GamePublisher) error
	CreateMultiple(ctx context.Context, items []entity.GamePublisher) error
	Delete(ctx context.Context, i *entity.GamePublisher) error
	DeleteMultiple(ctx context.Context, items []entity.GamePublisher) error

	FindByPublisherID(ctx context.Context, publisherID uint) ([]entity.GamePublisher, error)
	FindByGameID(ctx context.Context, gameID uint) ([]entity.GamePublisher, error)
	FindByGameIDAndPublisherIDs(ctx context.Context, gameID uint, publisherIDs []uint) ([]entity.GamePublisher, error)
}
