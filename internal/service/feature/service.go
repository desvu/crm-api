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

	err = s.GameFeatureRepository.DeleteMultiple(ctx, s.getGameFeaturesForDelete(featureIDs, currentGameFeatures))
	if err != nil {
		return err
	}

	err = s.GameFeatureRepository.CreateMultiple(ctx, s.getGameFeaturesForInsert(game.ID, featureIDs, currentGameFeatures))
	if err != nil {
		return err
	}

	return nil
}

func (s Service) getGameFeaturesForInsert(gameID uint, newFeatureIDs []uint, currentGameFeatures []entity.GameFeature) []entity.GameFeature {
	gameFeatures := make([]entity.GameFeature, len(newFeatureIDs))
	for i := range newFeatureIDs {
		gameFeatures[i] = entity.GameFeature{
			GameID:    gameID,
			FeatureID: newFeatureIDs[i],
		}
	}

	for i := 0; i < len(gameFeatures); i++ {
		var hasMatch bool
		for j := range currentGameFeatures {
			if gameFeatures[i].FeatureID == currentGameFeatures[j].FeatureID {
				hasMatch = true
			}
		}

		if hasMatch {
			gameFeatures = append(gameFeatures[:i], gameFeatures[i+1:]...)
			i--
		}
	}

	return gameFeatures
}

func (s Service) getGameFeaturesForDelete(newFeatureIDs []uint, currentGameFeatures []entity.GameFeature) []entity.GameFeature {
	gameFeatures := currentGameFeatures
	for i := 0; i < len(gameFeatures); i++ {
		var hasMatch bool
		for j := range newFeatureIDs {
			if gameFeatures[i].FeatureID == newFeatureIDs[j] {
				hasMatch = true
			}
		}

		if hasMatch {
			gameFeatures = append(gameFeatures[:i], gameFeatures[i+1:]...)
			i--
		}
	}

	return gameFeatures
}
