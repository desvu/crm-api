package postgres

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type GameFeatureRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) GameFeatureRepository {
	return GameFeatureRepository{
		h: env.Handler,
	}
}

func (r GameFeatureRepository) Create(ctx context.Context, i *entity.GameFeature) error {
	model, err := newModel(i)
	if err != nil {
		return err
	}

	_, err = r.h.ModelContext(ctx, model).Insert()
	if err != nil {
		return err
	}

	*i = *model.Convert()
	return nil
}

func (r GameFeatureRepository) CreateMultiple(ctx context.Context, items []entity.GameFeature) error {
	models := make([]model, len(items))
	for i := range items {
		m, err := newModel(&items[i])
		if err != nil {
			return err
		}

		models[i] = *m
	}

	_, err := r.h.ModelContext(ctx, models).Insert()
	if err != nil {
		return err
	}

	for i := range models {
		items[i] = *models[i].Convert()
	}

	return nil
}

func (r GameFeatureRepository) Delete(ctx context.Context, i *entity.GameFeature) error {
	model, err := newModel(i)
	if err != nil {
		return err
	}

	_, err = r.h.ModelContext(ctx, model).WherePK().Delete()
	if err != nil {
		return err
	}

	*i = *model.Convert()
	return nil
}

func (r GameFeatureRepository) DeleteMultiple(ctx context.Context, items []entity.GameFeature) error {
	models := make([]model, len(items))
	for i := range items {
		m, err := newModel(&items[i])
		if err != nil {
			return err
		}

		models[i] = *m
	}

	_, err := r.h.ModelContext(ctx, models).Delete()
	if err != nil {
		return err
	}

	return nil
}

func (r GameFeatureRepository) FindByGameID(ctx context.Context, gameID uint) ([]entity.GameFeature, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("game_id = ?", gameID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameFeature, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameFeatureRepository) FindByFeatureID(ctx context.Context, featureID uint) ([]entity.GameFeature, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("feature_id = ?", featureID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameFeature, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameFeatureRepository) FindByGameIDAndFeatureIDs(ctx context.Context, gameID uint, featureIDs []uint) ([]entity.GameFeature, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).
		Where("game_id = ?", gameID).
		Where("feature_id in (?)", featureIDs).
		Select()

	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameFeature, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
