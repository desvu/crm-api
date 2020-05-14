package postgres

import (
	"context"

	"github.com/pkg/errors"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type GameRevisionFeatureRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) GameRevisionFeatureRepository {
	return GameRevisionFeatureRepository{
		h: env.Handler,
	}
}

func (r GameRevisionFeatureRepository) Create(ctx context.Context, i *entity.GameRevisionFeature) error {
	model, err := newModel(i)
	if err != nil {
		return err
	}

	_, err = r.h.ModelContext(ctx, model).Insert()
	if err != nil {
		return errors.Wrap(err, "insert feature failed")
	}

	*i = *model.Convert()
	return nil
}

func (r GameRevisionFeatureRepository) CreateMultiple(ctx context.Context, items []entity.GameRevisionFeature) error {
	if len(items) == 0 {
		return nil
	}

	models := make([]model, len(items))
	for i := range items {
		m, err := newModel(&items[i])
		if err != nil {
			return err
		}
		models[i] = *m
	}

	_, err := r.h.ModelContext(ctx, &models).Insert()
	if err != nil {
		return errors.Wrap(err, "insert features failed")
	}

	for i := range models {
		items[i] = *models[i].Convert()
	}

	return nil
}

func (r GameRevisionFeatureRepository) Delete(ctx context.Context, i *entity.GameRevisionFeature) error {
	model, err := newModel(i)
	if err != nil {
		return err
	}

	_, err = r.h.ModelContext(ctx, model).WherePK().Delete()
	if err != nil {
		return errors.Wrap(err, "delete feature failed")
	}

	*i = *model.Convert()
	return nil
}

func (r GameRevisionFeatureRepository) DeleteMultiple(ctx context.Context, items []entity.GameRevisionFeature) error {
	if len(items) == 0 {
		return nil
	}

	models := make([]model, len(items))
	for i := range items {
		m, err := newModel(&items[i])
		if err != nil {
			return err
		}

		models[i] = *m
	}

	_, err := r.h.ModelContext(ctx, &models).Delete()
	if err != nil {
		return errors.Wrap(err, "delete features failed")
	}

	return nil
}

func (r GameRevisionFeatureRepository) FindByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.GameRevisionFeature, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("game_revision_id = ?", gameRevisionID).Select()
	if err != nil {
		return nil, errors.Wrap(err, "load features failed")
	}

	entities := make([]entity.GameRevisionFeature, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionFeatureRepository) FindByGameRevisionIDs(ctx context.Context, gameRevisionIDs []uint) ([]entity.GameRevisionFeature, error) {
	if len(gameRevisionIDs) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).WhereIn("game_revision_id in (?)", gameRevisionIDs).Select()
	if err != nil {
		return nil, errors.Wrap(err, "load features failed")
	}

	entities := make([]entity.GameRevisionFeature, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionFeatureRepository) FindByFeatureID(ctx context.Context, featureID uint) ([]entity.GameRevisionFeature, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("feature_id = ?", featureID).Select()
	if err != nil {
		return nil, errors.Wrap(err, "load features failed")
	}

	entities := make([]entity.GameRevisionFeature, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionFeatureRepository) FindByGameRevisionIDAndFeatureIDs(ctx context.Context, gameRevisionID uint, featureIDs []uint) ([]entity.GameRevisionFeature, error) {
	if len(featureIDs) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).
		Where("game_revision_id = ?", gameRevisionID).
		WhereIn("feature_id in (?)", featureIDs).
		Select()

	if err != nil {
		return nil, errors.Wrap(err, "load features failed")
	}

	entities := make([]entity.GameRevisionFeature, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
