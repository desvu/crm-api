package developer

import (
	"context"
	"errors"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

var ErrDeveloperNotFound = errors.New("developer not found")
var ErrInvalidDeveloperIDs = errors.New("invalid developer ids")

func (s Service) Create(ctx context.Context, data *service.CreateDeveloperData) (*entity.Developer, error) {
	developer := &entity.Developer{
		Name: data.Name,
	}

	if err := s.DeveloperRepository.Create(ctx, developer); err != nil {
		return nil, err
	}

	return developer, nil
}

func (s Service) Update(ctx context.Context, data *service.UpdateDeveloperData) (*entity.Developer, error) {
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

func (s Service) Delete(ctx context.Context, id uint) error {
	developer, err := s.GetExistByID(ctx, id)
	if err != nil {
		return err
	}

	return s.DeveloperRepository.Delete(ctx, developer)
}

func (s Service) GetByID(ctx context.Context, id uint) (*entity.Developer, error) {
	return s.DeveloperRepository.FindByID(ctx, id)
}

func (s Service) GetExistByID(ctx context.Context, id uint) (*entity.Developer, error) {
	developer, err := s.DeveloperRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if developer == nil {
		return nil, ErrDeveloperNotFound
	}

	return developer, nil
}

func (s Service) GetByIDs(ctx context.Context, ids []uint) ([]entity.Developer, error) {
	return s.DeveloperRepository.FindByIDs(ctx, ids)
}

func (s Service) GetByGameID(ctx context.Context, gameID uint) ([]entity.Developer, error) {
	gameDevelopers, err := s.GameDeveloperRepository.FindByGameID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	return s.GetByIDs(ctx, entity.NewGameDeveloperArray(gameDevelopers).IDs())
}

func (s Service) UpdateDevelopersForGame(ctx context.Context, game *entity.Game, developerIDs []uint) error {
	developers, err := s.GetByIDs(ctx, developerIDs)
	if err != nil {
		return err
	}

	// checking for IDs among the developers
	if len(developers) != len(developerIDs) {
		return ErrInvalidDeveloperIDs
	}

	currentGameDevelopers, err := s.GameDeveloperRepository.FindByGameID(ctx, game.ID)
	if err != nil {
		return err
	}

	err = s.GameDeveloperRepository.DeleteMultiple(ctx, s.getGameDevelopersForDelete(developerIDs, currentGameDevelopers))
	if err != nil {
		return err
	}

	err = s.GameDeveloperRepository.CreateMultiple(ctx, s.getGameDevelopersForInsert(game.ID, developerIDs, currentGameDevelopers))
	if err != nil {
		return err
	}

	return nil
}

func (s Service) getGameDevelopersForInsert(gameID uint, newDeveloperIDs []uint, currentGameDevelopers []entity.GameDeveloper) []entity.GameDeveloper {
	gameDevelopers := make([]entity.GameDeveloper, len(newDeveloperIDs))
	for i := range newDeveloperIDs {
		gameDevelopers[i] = entity.GameDeveloper{
			GameID:      gameID,
			DeveloperID: newDeveloperIDs[i],
		}
	}

	for i := 0; i < len(gameDevelopers); i++ {
		var hasMatch bool
		for j := range currentGameDevelopers {
			if gameDevelopers[i].DeveloperID == currentGameDevelopers[j].DeveloperID {
				hasMatch = true
			}
		}

		if hasMatch {
			gameDevelopers = append(gameDevelopers[:i], gameDevelopers[i+1:]...)
			i--
		}
	}

	return gameDevelopers
}

func (s Service) getGameDevelopersForDelete(newDeveloperIDs []uint, currentGameDevelopers []entity.GameDeveloper) []entity.GameDeveloper {
	gameDevelopers := currentGameDevelopers
	for i := 0; i < len(gameDevelopers); i++ {
		var hasMatch bool
		for j := range newDeveloperIDs {
			if gameDevelopers[i].DeveloperID == newDeveloperIDs[j] {
				hasMatch = true
			}
		}

		if hasMatch {
			gameDevelopers = append(gameDevelopers[:i], gameDevelopers[i+1:]...)
			i--
		}
	}

	return gameDevelopers
}
