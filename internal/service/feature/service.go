package feature

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/service"
	"github.com/qilin/crm-api/pkg/errors"
)

type Service struct {
	ServiceParams
}

var ErrFeatureNotFound = errors.NewService(errors.ErrNotFound, "feature not found")
var ErrInvalidFeatureIDs = errors.NewService(errors.ErrValidation, "invalid feature ids")

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

func (s Service) GetByGameRevisionID(ctx context.Context, gameID uint) ([]entity.Feature, error) {
	gameFeatures, err := s.GameRevisionFeatureRepository.FindByGameRevisionID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	return s.GetByIDs(ctx, entity.NewGameRevisionFeatureArray(gameFeatures).IDs())
}

func (s Service) UpdateFeaturesForGameRevision(ctx context.Context, gameRevision *entity.GameRevision, featureIDs []uint) error {
	features, err := s.GetByIDs(ctx, featureIDs)
	if err != nil {
		return err
	}

	// checking for IDs among the features
	if len(features) != len(featureIDs) {
		return ErrInvalidFeatureIDs
	}

	currentGameFeatures, err := s.GameRevisionFeatureRepository.FindByGameRevisionID(ctx, gameRevision.ID)
	if err != nil {
		return err
	}

	err = s.GameRevisionFeatureRepository.DeleteMultiple(ctx, getGameFeaturesForDelete(featureIDs, currentGameFeatures))
	if err != nil {
		return err
	}

	err = s.GameRevisionFeatureRepository.CreateMultiple(ctx, getGameFeaturesForInsert(gameRevision.ID, featureIDs, currentGameFeatures))
	if err != nil {
		return err
	}

	return nil
}

func getGameFeaturesForInsert(gameID uint, newFeatureIDs []uint, currentGameFeatures []entity.GameRevisionFeature) []entity.GameRevisionFeature {
	gameFeatures := make([]entity.GameRevisionFeature, 0)
	for _, newFeatureID := range newFeatureIDs {
		var hasMatch bool
		for _, currentGameFeature := range currentGameFeatures {
			if newFeatureID == currentGameFeature.FeatureID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameFeatures = append(gameFeatures, entity.GameRevisionFeature{
				GameRevisionID: gameID,
				FeatureID:      newFeatureID,
			})
		}
	}

	return gameFeatures
}

func getGameFeaturesForDelete(newFeatureIDs []uint, currentGameFeatures []entity.GameRevisionFeature) []entity.GameRevisionFeature {
	gameFeatures := make([]entity.GameRevisionFeature, 0)
	for _, currentGameFeature := range currentGameFeatures {
		var hasMatch bool
		for _, newFeatureID := range newFeatureIDs {
			if currentGameFeature.FeatureID == newFeatureID {
				hasMatch = true
			}
		}

		if !hasMatch {
			gameFeatures = append(gameFeatures, entity.GameRevisionFeature{
				ID:             currentGameFeature.ID,
				GameRevisionID: currentGameFeature.GameRevisionID,
				FeatureID:      currentGameFeature.FeatureID,
			})
		}
	}

	return gameFeatures
}
