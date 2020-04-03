package postgres

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
)

type GameDeveloperRepository struct {
	db *pg.DB
}

func New(env *env.Postgres) GameDeveloperRepository {
	return GameDeveloperRepository{db: env.DB}
}

func (r GameDeveloperRepository) Create(ctx context.Context, i *entity.GameDeveloper) error {
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

func (r GameDeveloperRepository) CreateMultiple(ctx context.Context, items []entity.GameDeveloper) error {
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

func (r GameDeveloperRepository) Delete(ctx context.Context, i *entity.GameDeveloper) error {
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

func (r GameDeveloperRepository) DeleteMultiple(ctx context.Context, items []entity.GameDeveloper) error {
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

func (r GameDeveloperRepository) FindByGameID(ctx context.Context, gameID uint) ([]entity.GameDeveloper, error) {
	var models []model

	err := r.db.ModelContext(ctx, &models).Where("game_id = ?", gameID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameDeveloper, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameDeveloperRepository) FindByDeveloperID(ctx context.Context, developerID uint) ([]entity.GameDeveloper, error) {
	var models []model

	err := r.db.ModelContext(ctx, &models).Where("developer_id = ?", developerID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameDeveloper, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameDeveloperRepository) FindByGameIDAndDeveloperIDs(ctx context.Context, gameID uint, developerIDs []uint) ([]entity.GameDeveloper, error) {
	var models []model

	err := r.db.ModelContext(ctx, &models).
		Where("game_id = ?", gameID).
		Where("developer_id in (?)", developerIDs).
		Select()

	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameDeveloper, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
