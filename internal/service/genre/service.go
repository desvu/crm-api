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

func (s Service) UpdateGenreForGame(ctx context.Context, game *entity.Game, genreIDs []uint) error {
	genres, err := s.GetByIDs(ctx, genreIDs)
	if err != nil {
		return err
	}

	// checking for IDs among the genres
	if len(genres) != len(genreIDs) {
		return ErrInvalidGenreIDs
	}

	currentGameGenre, err := s.GameGenreRepository.FindByGameID(ctx, game.ID)
	if err != nil {
		return err
	}

	err = s.GameGenreRepository.DeleteMultiple(ctx, s.getGameGenreForDelete(genreIDs, currentGameGenre))
	if err != nil {
		return err
	}

	err = s.GameGenreRepository.CreateMultiple(ctx, s.getGameGenreForInsert(game.ID, genreIDs, currentGameGenre))
	if err != nil {
		return err
	}

	return nil
}

func (s Service) getGameGenreForInsert(gameID uint, newGenreIDs []uint, currentGameGenre []entity.GameGenre) []entity.GameGenre {
	gameGenre := make([]entity.GameGenre, len(newGenreIDs))
	for i := range newGenreIDs {
		gameGenre[i] = entity.GameGenre{
			GameID:  gameID,
			GenreID: newGenreIDs[i],
		}
	}

	for i := 0; i < len(gameGenre); i++ {
		var hasMatch bool
		for j := range currentGameGenre {
			if gameGenre[i].GenreID == currentGameGenre[j].GenreID {
				hasMatch = true
			}
		}

		if hasMatch {
			gameGenre = append(gameGenre[:i], gameGenre[i+1:]...)
			i--
		}
	}

	return gameGenre
}

func (s Service) getGameGenreForDelete(newGenreIDs []uint, currentGameGenre []entity.GameGenre) []entity.GameGenre {
	gameGenre := currentGameGenre
	for i := 0; i < len(gameGenre); i++ {
		var hasMatch bool
		for j := range newGenreIDs {
			if gameGenre[i].GenreID == newGenreIDs[j] {
				hasMatch = true
			}
		}

		if hasMatch {
			gameGenre = append(gameGenre[:i], gameGenre[i+1:]...)
			i--
		}
	}

	return gameGenre
}
