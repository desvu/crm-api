package postgres

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
)

type GameTagRepository struct {
	db *pg.DB
}

func New(env *env.Postgres) GameTagRepository {
	return GameTagRepository{db: env.DB}
}

func (r GameTagRepository) Create(ctx context.Context, i *entity.GameTag) error {
	model, err := newModel(i)
	if err != nil {
		return err
	}

	_, err = r.db.ModelContext(ctx, model).Insert()
	if err != nil {
		return err
	}

	*i = *model.Convert()
	return nil
}

func (r GameTagRepository) CreateMultiple(ctx context.Context, items []entity.GameTag) error {
	models := make([]model, len(items))
	for i := range items {
		m, err := newModel(&items[i])
		if err != nil {
			return err
		}

		models[i] = *m
	}

	_, err := r.db.ModelContext(ctx, models).Insert()
	if err != nil {
		return err
	}

	for i := range models {
		items[i] = *models[i].Convert()
	}

	return nil
}

func (r GameTagRepository) Delete(ctx context.Context, i *entity.GameTag) error {
	model, err := newModel(i)
	if err != nil {
		return err
	}

	_, err = r.db.ModelContext(ctx, model).WherePK().Delete()
	if err != nil {
		return err
	}

	*i = *model.Convert()
	return nil
}

func (r GameTagRepository) DeleteMultiple(ctx context.Context, items []entity.GameTag) error {
	models := make([]model, len(items))
	for i := range items {
		m, err := newModel(&items[i])
		if err != nil {
			return err
		}

		models[i] = *m
	}

	_, err := r.db.ModelContext(ctx, models).Delete()
	if err != nil {
		return err
	}

	return nil
}

func (r GameTagRepository) FindByGameID(ctx context.Context, gameID uint) ([]entity.GameTag, error) {
	var models []model

	err := r.db.ModelContext(ctx, &models).Where("game_id = ?", gameID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameTag, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameTagRepository) FindByTagID(ctx context.Context, tagID uint) ([]entity.GameTag, error) {
	var models []model

	err := r.db.ModelContext(ctx, &models).Where("tag_id = ?", tagID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameTag, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameTagRepository) FindByGameIDAndTagIDs(ctx context.Context, gameID uint, tagIDs []uint) ([]entity.GameTag, error) {
	var models []model

	err := r.db.ModelContext(ctx, &models).
		Where("game_id = ?", gameID).
		Where("tag_id in (?)", tagIDs).
		Select()

	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameTag, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
