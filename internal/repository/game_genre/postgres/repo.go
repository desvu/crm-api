package postgres

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
)

type GameGenreRepository struct {
	db *pg.DB
}

func New(env *env.Postgres) GameGenreRepository {
	return GameGenreRepository{db: env.DB}
}

func (r GameGenreRepository) Create(ctx context.Context, i *entity.GameGenre) error {
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

func (r GameGenreRepository) CreateMultiple(ctx context.Context, items []entity.GameGenre) error {
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

func (r GameGenreRepository) Delete(ctx context.Context, i *entity.GameGenre) error {
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

func (r GameGenreRepository) DeleteMultiple(ctx context.Context, items []entity.GameGenre) error {
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

func (r GameGenreRepository) FindByGameID(ctx context.Context, gameID uint) ([]entity.GameGenre, error) {
	var models []model

	err := r.db.ModelContext(ctx, &models).Where("game_id = ?", gameID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameGenre, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameGenreRepository) FindByGenreID(ctx context.Context, genreID uint) ([]entity.GameGenre, error) {
	var models []model

	err := r.db.ModelContext(ctx, &models).Where("genre_id = ?", genreID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameGenre, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameGenreRepository) FindByGameIDAndGenreIDs(ctx context.Context, gameID uint, genreIDs []uint) ([]entity.GameGenre, error) {
	var models []model

	err := r.db.ModelContext(ctx, &models).
		Where("game_id = ?", gameID).
		Where("genre_id in (?)", genreIDs).
		Select()

	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameGenre, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
