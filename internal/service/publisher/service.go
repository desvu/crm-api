package publisher

import (
	"context"
	"errors"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

var ErrPublisherNotFound = errors.New("publisher not found")
var ErrInvalidPublisherIDs = errors.New("invalid publisher ids")

func (s Service) Create(ctx context.Context, data *service.CreatePublisherData) (*entity.Publisher, error) {
	publisher := &entity.Publisher{
		Name: data.Name,
	}

	if err := s.PublisherRepository.Create(ctx, publisher); err != nil {
		return nil, err
	}

	return publisher, nil
}

func (s Service) Update(ctx context.Context, data *service.UpdatePublisherData) (*entity.Publisher, error) {
	publisher, err := s.GetExistByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if publisher.Name != data.Name {
		publisher.Name = data.Name
		if err = s.PublisherRepository.Update(ctx, publisher); err != nil {
			return nil, err
		}
	}

	return publisher, nil
}

func (s Service) Delete(ctx context.Context, id uint) error {
	publisher, err := s.GetExistByID(ctx, id)
	if err != nil {
		return err
	}

	return s.PublisherRepository.Delete(ctx, publisher)
}

func (s Service) GetByID(ctx context.Context, id uint) (*entity.Publisher, error) {
	return s.PublisherRepository.FindByID(ctx, id)
}

func (s Service) GetExistByID(ctx context.Context, id uint) (*entity.Publisher, error) {
	publisher, err := s.PublisherRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if publisher == nil {
		return nil, ErrPublisherNotFound
	}

	return publisher, nil
}

func (s Service) GetByIDs(ctx context.Context, ids []uint) ([]entity.Publisher, error) {
	return s.PublisherRepository.FindByIDs(ctx, ids)
}

func (s Service) GetByGameID(ctx context.Context, gameID uint) ([]entity.Publisher, error) {
	gamePublishers, err := s.GamePublisherRepository.FindByGameID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	return s.GetByIDs(ctx, entity.NewGamePublisherArray(gamePublishers).IDs())
}

func (s Service) UpdatePublishersForGame(ctx context.Context, game *entity.Game, publisherIDs []uint) error {
	publishers, err := s.GetByIDs(ctx, publisherIDs)
	if err != nil {
		return err
	}

	// checking for IDs among the publishers
	if len(publishers) != len(publisherIDs) {
		return ErrInvalidPublisherIDs
	}

	currentGamePublisher, err := s.GamePublisherRepository.FindByGameID(ctx, game.ID)
	if err != nil {
		return err
	}

	err = s.GamePublisherRepository.DeleteMultiple(ctx, getGamePublishersForDelete(publisherIDs, currentGamePublisher))
	if err != nil {
		return err
	}

	err = s.GamePublisherRepository.CreateMultiple(ctx, getGamePublishersForInsert(game.ID, publisherIDs, currentGamePublisher))
	if err != nil {
		return err
	}

	return nil
}

func getGamePublishersForInsert(gameID uint, newPublisherIDs []uint, currentGamePublisher []entity.GamePublisher) []entity.GamePublisher {
	gamePublisher := make([]entity.GamePublisher, 0)
	for _, newPublisherID := range newPublisherIDs {
		var hasMatch bool
		for _, currentGamePublisher := range currentGamePublisher {
			if newPublisherID == currentGamePublisher.PublisherID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gamePublisher = append(gamePublisher, entity.GamePublisher{
				GameID:      gameID,
				PublisherID: newPublisherID,
			})
		}
	}

	return gamePublisher
}

func getGamePublishersForDelete(newPublisherIDs []uint, currentGamePublisher []entity.GamePublisher) []entity.GamePublisher {
	gamePublisher := make([]entity.GamePublisher, 0)
	for _, currentGamePublisher := range currentGamePublisher {
		var hasMatch bool
		for _, newPublisherID := range newPublisherIDs {
			if currentGamePublisher.PublisherID == newPublisherID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gamePublisher = append(gamePublisher, entity.GamePublisher{
				ID:          currentGamePublisher.ID,
				GameID:      currentGamePublisher.GameID,
				PublisherID: currentGamePublisher.PublisherID,
			})
		}
	}

	return gamePublisher
}
