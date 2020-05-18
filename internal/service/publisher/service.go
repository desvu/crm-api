package publisher

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/errors"
	"github.com/qilin/crm-api/internal/domain/repository"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

func (s *Service) Create(ctx context.Context, data *service.CreatePublisherData) (*entity.Publisher, error) {
	publisher := &entity.Publisher{
		Name: data.Name,
	}

	if err := s.PublisherRepository.Create(ctx, publisher); err != nil {
		return nil, err
	}

	return publisher, nil
}

func (s *Service) Update(ctx context.Context, data *service.UpdatePublisherData) (*entity.Publisher, error) {
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

func (s *Service) Delete(ctx context.Context, id uint) error {
	publisher, err := s.GetExistByID(ctx, id)
	if err != nil {
		return err
	}

	return s.PublisherRepository.Delete(ctx, publisher)
}

func (s *Service) GetByID(ctx context.Context, id uint) (*entity.Publisher, error) {
	return s.PublisherRepository.FindByID(ctx, id)
}

func (s *Service) GetExistByID(ctx context.Context, id uint) (*entity.Publisher, error) {
	publisher, err := s.PublisherRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if publisher == nil {
		return nil, errors.PublisherNotFound
	}

	return publisher, nil
}

func (s *Service) GetByIDs(ctx context.Context, ids []uint) ([]entity.Publisher, error) {
	return s.PublisherRepository.FindByIDs(ctx, ids)
}

func (s *Service) GetByGameRevisionID(ctx context.Context, gameID uint) ([]entity.Publisher, error) {
	gamePublishers, err := s.GameRevisionPublisherRepository.FindByGameRevisionID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	return s.GetByIDs(ctx, entity.NewGameRevisionPublisherArray(gamePublishers).IDs())
}

func (s *Service) GetByFilter(ctx context.Context, data *service.GetByFilterPublisherData) ([]entity.Publisher, error) {
	return s.PublisherRepository.FindByFilter(ctx, &repository.FindByFilterPublisherData{
		Limit:  data.Limit,
		Offset: data.Offset,
	})
}

func (s *Service) UpdatePublishersForGameRevision(ctx context.Context, gameRevision *entity.GameRevision, publisherIDs []uint) error {
	publishers, err := s.GetByIDs(ctx, publisherIDs)
	if err != nil {
		return err
	}

	// checking for IDs among the publishers
	if len(publishers) != len(publisherIDs) {
		return errors.InvalidPublisherIDs
	}

	currentGamePublisher, err := s.GameRevisionPublisherRepository.FindByGameRevisionID(ctx, gameRevision.ID)
	if err != nil {
		return err
	}

	err = s.GameRevisionPublisherRepository.DeleteMultiple(ctx, getGamePublishersForDelete(publisherIDs, currentGamePublisher))
	if err != nil {
		return err
	}

	err = s.GameRevisionPublisherRepository.CreateMultiple(ctx, getGamePublishersForInsert(gameRevision.ID, publisherIDs, currentGamePublisher))
	if err != nil {
		return err
	}

	return nil
}

func getGamePublishersForInsert(gameID uint, newPublisherIDs []uint, currentGamePublisher []entity.GameRevisionPublisher) []entity.GameRevisionPublisher {
	gamePublisher := make([]entity.GameRevisionPublisher, 0)
	for _, newPublisherID := range newPublisherIDs {
		var hasMatch bool
		for _, currentGamePublisher := range currentGamePublisher {
			if newPublisherID == currentGamePublisher.PublisherID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gamePublisher = append(gamePublisher, entity.GameRevisionPublisher{
				GameRevisionID: gameID,
				PublisherID:    newPublisherID,
			})
		}
	}

	return gamePublisher
}

func getGamePublishersForDelete(newPublisherIDs []uint, currentGamePublisher []entity.GameRevisionPublisher) []entity.GameRevisionPublisher {
	gamePublisher := make([]entity.GameRevisionPublisher, 0)
	for _, currentGamePublisher := range currentGamePublisher {
		var hasMatch bool
		for _, newPublisherID := range newPublisherIDs {
			if currentGamePublisher.PublisherID == newPublisherID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gamePublisher = append(gamePublisher, entity.GameRevisionPublisher{
				ID:             currentGamePublisher.ID,
				GameRevisionID: currentGamePublisher.GameRevisionID,
				PublisherID:    currentGamePublisher.PublisherID,
			})
		}
	}

	return gamePublisher
}
