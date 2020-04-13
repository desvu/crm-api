package feature

import (
	"context"
	"errors"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/service"
)

type Service struct {
	ServiceParams
}

var ErrFeatureNotFound = errors.New("feature not found")
var ErrInvalidFeatureIDs = errors.New("invalid feature ids")

func (s Service) Create(ctx context.Context, data *service.CreateFeatureData) (*entity.Feature, error) {
	feature := &entity.Feature{
		Name: data.Name,
	}

	if err := s.FeatureRepository.Create(ctx, feature); err != nil {
		return nil, err
	}

	return feature, nil
}

func (s Service) Update(ctx context.Context, data *service.UpdateFeatureData) (*entity.Feature, error) {
	feature, err := s.GetExistByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if feature.Name != data.Name {
		feature.Name = data.Name
		if err = s.FeatureRepository.Update(ctx, feature); err != nil {
			return nil, err
		}
	}

	return feature, nil
}

func (s Service) Delete(ctx context.Context, id uint) error {
	feature, err := s.GetExistByID(ctx, id)
	if err != nil {
		return err
	}

	return s.FeatureRepository.Delete(ctx, feature)
}

func (s Service) GetByID(ctx context.Context, id uint) (*entity.Feature, error) {
	return s.FeatureRepository.FindByID(ctx, id)
}

func (s Service) GetExistByID(ctx context.Context, id uint) (*entity.Feature, error) {
	feature, err := s.FeatureRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if feature == nil {
		return nil, ErrFeatureNotFound
	}

	return feature, nil
}

func (s Service) GetByIDs(ctx context.Context, ids []uint) ([]entity.Feature, error) {
	return s.FeatureRepository.FindByIDs(ctx, ids)
}

func (s Service) GetByGameID(ctx context.Context, gameID uint) ([]entity.Feature, error) {
	gameFeatures, err := s.GameFeatureRepository.FindByGameID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	return s.GetByIDs(ctx, entity.NewGameFeatureArray(gameFeatures).IDs())
}

func (s Service) UpdateFeaturesForGame(ctx context.Context, game *entity.Game, featureIDs []uint) error {
	features, err := s.GetByIDs(ctx, featureIDs)
	if err != nil {
		return err
	}

	// checking for IDs among the features
	if len(features) != len(featureIDs) {
		return ErrInvalidFeatureIDs
	}

	currentGameFeatures, err := s.GameFeatureRepository.FindByGameID(ctx, game.ID)
	if err != nil {
		return err
	}

	err = s.GameFeatureRepository.DeleteMultiple(ctx, getGameFeaturesForDelete(featureIDs, currentGameFeatures))
	if err != nil {
		return err
	}

	err = s.GameFeatureRepository.CreateMultiple(ctx, getGameFeaturesForInsert(game.ID, featureIDs, currentGameFeatures))
	if err != nil {
		return err
	}

	return nil
}

func getGameFeaturesForInsert(gameID uint, newFeatureIDs []uint, currentGameFeatures []entity.GameFeature) []entity.GameFeature {
	gameFeatures := make([]entity.GameFeature, 0)
	for _, newFeatureID := range newFeatureIDs {
		var hasMatch bool
		for _, currentGameFeature := range currentGameFeatures {
			if newFeatureID == currentGameFeature.FeatureID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameFeatures = append(gameFeatures, entity.GameFeature{
				GameID:    gameID,
				FeatureID: newFeatureID,
			})
		}
	}

	return gameFeatures
}

func getGameFeaturesForDelete(newFeatureIDs []uint, currentGameFeatures []entity.GameFeature) []entity.GameFeature {
	gameFeatures := make([]entity.GameFeature, 0)
	for _, currentGameFeature := range currentGameFeatures {
		var hasMatch bool
		for _, newFeatureID := range newFeatureIDs {
			if currentGameFeature.FeatureID == newFeatureID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameFeatures = append(gameFeatures, entity.GameFeature{
				ID:        currentGameFeature.ID,
				GameID:    currentGameFeature.GameID,
				FeatureID: currentGameFeature.FeatureID,
			})
		}
	}

	return gameFeatures
}
