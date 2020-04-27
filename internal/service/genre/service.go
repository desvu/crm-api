package genre

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/errors"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

func (s Service) Create(ctx context.Context, data *service.CreateGenreData) (*entity.Genre, error) {
	genre := &entity.Genre{
		Name: data.Name,
	}

	if err := s.GenreRepository.Create(ctx, genre); err != nil {
		return nil, err
	}

	return genre, nil
}

func (s Service) Update(ctx context.Context, data *service.UpdateGenreData) (*entity.Genre, error) {
	genre, err := s.GetExistByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if genre.Name != data.Name {
		genre.Name = data.Name
		if err = s.GenreRepository.Update(ctx, genre); err != nil {
			return nil, err
		}
	}

	return genre, nil
}

func (s Service) Delete(ctx context.Context, id uint) error {
	genre, err := s.GetExistByID(ctx, id)
	if err != nil {
		return err
	}

	return s.GenreRepository.Delete(ctx, genre)
}

func (s Service) GetByID(ctx context.Context, id uint) (*entity.Genre, error) {
	return s.GenreRepository.FindByID(ctx, id)
}

func (s Service) GetExistByID(ctx context.Context, id uint) (*entity.Genre, error) {
	genre, err := s.GenreRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if genre == nil {
		return nil, errors.GenreNotFound
	}

	return genre, nil
}

func (s Service) GetByIDs(ctx context.Context, ids []uint) ([]entity.Genre, error) {
	return s.GenreRepository.FindByIDs(ctx, ids)
}

func (s Service) GetByGameRevisionID(ctx context.Context, gameID uint) ([]entity.Genre, error) {
	gameGenres, err := s.GameRevisionGenreRepository.FindByGameRevisionID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	return s.GetByIDs(ctx, entity.NewGameRevisionGenreArray(gameGenres).IDs())
}

func (s Service) UpdateGenresForGameRevision(ctx context.Context, gameRevision *entity.GameRevision, genreIDs []uint) error {
	genres, err := s.GetByIDs(ctx, genreIDs)
	if err != nil {
		return err
	}

	// checking for IDs among the genres
	if len(genres) != len(genreIDs) {
		return errors.InvalidGenreIDs
	}

	currentGameGenres, err := s.GameRevisionGenreRepository.FindByGameRevisionID(ctx, gameRevision.ID)
	if err != nil {
		return err
	}

	err = s.GameRevisionGenreRepository.DeleteMultiple(ctx, getGameGenresForDelete(genreIDs, currentGameGenres))
	if err != nil {
		return err
	}

	err = s.GameRevisionGenreRepository.CreateMultiple(ctx, getGameGenresForInsert(gameRevision.ID, genreIDs, currentGameGenres))
	if err != nil {
		return err
	}

	return nil
}

func getGameGenresForInsert(gameID uint, newGenreIDs []uint, currentGameGenres []entity.GameRevisionGenre) []entity.GameRevisionGenre {
	gameGenres := make([]entity.GameRevisionGenre, 0)
	for _, newGenreID := range newGenreIDs {
		var hasMatch bool
		for _, currentGameGenre := range currentGameGenres {
			if newGenreID == currentGameGenre.GenreID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameGenres = append(gameGenres, entity.GameRevisionGenre{
				GameRevisionID: gameID,
				GenreID:        newGenreID,
			})
		}
	}

	return gameGenres
}

func getGameGenresForDelete(newGenreIDs []uint, currentGameGenres []entity.GameRevisionGenre) []entity.GameRevisionGenre {
	gameGenres := make([]entity.GameRevisionGenre, 0)
	for _, currentGameGenre := range currentGameGenres {
		var hasMatch bool
		for _, newGenreID := range newGenreIDs {
			if currentGameGenre.GenreID == newGenreID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameGenres = append(gameGenres, entity.GameRevisionGenre{
				ID:             currentGameGenre.ID,
				GameRevisionID: currentGameGenre.GameRevisionID,
				GenreID:        currentGameGenre.GenreID,
			})
		}
	}

	return gameGenres
}
