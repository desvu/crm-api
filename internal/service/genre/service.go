package genre

import (
	"context"
	"errors"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

var ErrGenreNotFound = errors.New("genre not found")
var ErrInvalidGenreIDs = errors.New("invalid genre ids")

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
		return nil, ErrGenreNotFound
	}

	return genre, nil
}

func (s Service) GetByIDs(ctx context.Context, ids []uint) ([]entity.Genre, error) {
	return s.GenreRepository.FindByIDs(ctx, ids)
}

func (s Service) GetByGameID(ctx context.Context, gameID uint) ([]entity.Genre, error) {
	gameGenres, err := s.GameGenreRepository.FindByGameID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	return s.GetByIDs(ctx, entity.NewGameGenreArray(gameGenres).IDs())
}

func (s Service) UpdateGenresForGame(ctx context.Context, game *entity.Game, genreIDs []uint) error {
	genres, err := s.GetByIDs(ctx, genreIDs)
	if err != nil {
		return err
	}

	// checking for IDs among the genres
	if len(genres) != len(genreIDs) {
		return ErrInvalidGenreIDs
	}

	currentGameGenres, err := s.GameGenreRepository.FindByGameID(ctx, game.ID)
	if err != nil {
		return err
	}

	err = s.GameGenreRepository.DeleteMultiple(ctx, s.getGameGenresForDelete(genreIDs, currentGameGenres))
	if err != nil {
		return err
	}

	err = s.GameGenreRepository.CreateMultiple(ctx, s.getGameGenresForInsert(game.ID, genreIDs, currentGameGenres))
	if err != nil {
		return err
	}

	return nil
}

func (s Service) getGameGenresForInsert(gameID uint, newGenreIDs []uint, currentGameGenres []entity.GameGenre) []entity.GameGenre {
	gameGenres := make([]entity.GameGenre, 0)
	for _, newGenreID := range newGenreIDs {
		var hasMatch bool
		for _, currentGameGenre := range currentGameGenres {
			if newGenreID == currentGameGenre.GenreID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameGenres = append(gameGenres, entity.GameGenre{
				GameID:  gameID,
				GenreID: newGenreID,
			})
		}
	}

	return gameGenres
}

func (s Service) getGameGenresForDelete(newGenreIDs []uint, currentGameGenres []entity.GameGenre) []entity.GameGenre {
	gameGenres := make([]entity.GameGenre, 0)
	for _, currentGameGenre := range currentGameGenres {
		var hasMatch bool
		for _, newGenreID := range newGenreIDs {
			if currentGameGenre.GenreID == newGenreID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameGenres = append(gameGenres, entity.GameGenre{
				ID:      currentGameGenre.ID,
				GameID:  currentGameGenre.GameID,
				GenreID: currentGameGenre.GenreID,
			})
		}
	}

	return gameGenres
}
