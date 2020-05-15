package developer

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/errors"
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

func (s *Service) Create(ctx context.Context, data *service.CreateDeveloperData) (*entity.Developer, error) {
	developer := &entity.Developer{
		Name: data.Name,
	}

	if err := s.DeveloperRepository.Create(ctx, developer); err != nil {
		return nil, err
	}

	return developer, nil
}

func (s *Service) Update(ctx context.Context, data *service.UpdateDeveloperData) (*entity.Developer, error) {
	developer, err := s.GetExistByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if developer.Name != data.Name {
		developer.Name = data.Name
		if err = s.DeveloperRepository.Update(ctx, developer); err != nil {
			return nil, err
		}
	}

	return developer, nil
}

func (s *Service) Delete(ctx context.Context, id uint) error {
	developer, err := s.GetExistByID(ctx, id)
	if err != nil {
		return err
	}

	return s.DeveloperRepository.Delete(ctx, developer)
}

func (s *Service) GetByID(ctx context.Context, id uint) (*entity.Developer, error) {
	return s.DeveloperRepository.FindByID(ctx, id)
}

func (s *Service) GetExistByID(ctx context.Context, id uint) (*entity.Developer, error) {
	developer, err := s.DeveloperRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if developer == nil {
		return nil, errors.DeveloperNotFound
	}

	return developer, nil
}

func (s *Service) GetByIDs(ctx context.Context, ids []uint) ([]entity.Developer, error) {
	return s.DeveloperRepository.FindByIDs(ctx, ids)
}

func (s *Service) GetByGameRevisionID(ctx context.Context, gameID uint) ([]entity.Developer, error) {
	gameDevelopers, err := s.GameRevisionDeveloperRepository.FindByGameRevisionID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	return s.GetByIDs(ctx, entity.NewGameRevisionDeveloperArray(gameDevelopers).IDs())
}

func (s *Service) GetByFilter(ctx context.Context, data *service.GetByFilterDeveloperData) ([]entity.Developer, error) {
	return s.DeveloperRepository.FindByFilter(ctx, &repository.FindByFilterDeveloperData{
		Limit:  data.Limit,
		Offset: data.Offset,
	})
}

func (s *Service) UpdateDevelopersForGameRevision(ctx context.Context, gameRevision *entity.GameRevision, developerIDs []uint) error {
	developers, err := s.GetByIDs(ctx, developerIDs)
	if err != nil {
		return err
	}

	// checking for IDs among the developers
	if len(developers) != len(developerIDs) {
		return errors.InvalidDeveloperIDs
	}

	currentGameDevelopers, err := s.GameRevisionDeveloperRepository.FindByGameRevisionID(ctx, gameRevision.ID)
	if err != nil {
		return err
	}

	err = s.GameRevisionDeveloperRepository.DeleteMultiple(ctx, getGameDevelopersForDelete(developerIDs, currentGameDevelopers))
	if err != nil {
		return err
	}

	err = s.GameRevisionDeveloperRepository.CreateMultiple(ctx, getGameDevelopersForInsert(gameRevision.ID, developerIDs, currentGameDevelopers))
	if err != nil {
		return err
	}

	return nil
}

func getGameDevelopersForInsert(gameID uint, newDeveloperIDs []uint, currentGameDevelopers []entity.GameRevisionDeveloper) []entity.GameRevisionDeveloper {
	gameDevelopers := make([]entity.GameRevisionDeveloper, 0)
	for _, newDeveloperID := range newDeveloperIDs {
		var hasMatch bool
		for _, currentGameDeveloper := range currentGameDevelopers {
			if newDeveloperID == currentGameDeveloper.DeveloperID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameDevelopers = append(gameDevelopers, entity.GameRevisionDeveloper{
				GameRevisionID: gameID,
				DeveloperID:    newDeveloperID,
			})
		}
	}

	return gameDevelopers
}

func getGameDevelopersForDelete(newDeveloperIDs []uint, currentGameDevelopers []entity.GameRevisionDeveloper) []entity.GameRevisionDeveloper {
	gameDevelopers := make([]entity.GameRevisionDeveloper, 0)
	for _, currentGameDeveloper := range currentGameDevelopers {
		var hasMatch bool
		for _, newDeveloperID := range newDeveloperIDs {
			if currentGameDeveloper.DeveloperID == newDeveloperID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameDevelopers = append(gameDevelopers, entity.GameRevisionDeveloper{
				ID:             currentGameDeveloper.ID,
				GameRevisionID: currentGameDeveloper.GameRevisionID,
				DeveloperID:    currentGameDeveloper.DeveloperID,
			})
		}
	}

	return gameDevelopers
}
