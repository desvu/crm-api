package postgres

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type GameRevisionGenreRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) GameRevisionGenreRepository {
	return GameRevisionGenreRepository{
		h: env.Handler,
	}
}

func (r GameRevisionGenreRepository) Create(ctx context.Context, i *entity.GameRevisionGenre) error {
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

func (r GameRevisionGenreRepository) CreateMultiple(ctx context.Context, items []entity.GameRevisionGenre) error {
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

	_, err := r.h.ModelContext(ctx, models).Insert()
	if err != nil {
		return err
	}

	for i := range models {
		items[i] = *models[i].Convert()
	}

	return nil
}

func (r GameRevisionGenreRepository) Delete(ctx context.Context, i *entity.GameRevisionGenre) error {
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

func (r GameRevisionGenreRepository) DeleteMultiple(ctx context.Context, items []entity.GameRevisionGenre) error {
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

	_, err := r.h.ModelContext(ctx, models).Delete()
	if err != nil {
		return err
	}

	return nil
}

func (r GameRevisionGenreRepository) FindByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.GameRevisionGenre, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("game_revision_id = ?", gameRevisionID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevisionGenre, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionGenreRepository) FindByGameRevisionIDs(ctx context.Context, gameRevisionIDs []uint) ([]entity.GameRevisionGenre, error) {
	if len(gameRevisionIDs) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).WhereIn("game_revision_id in (?)", gameRevisionIDs).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevisionGenre, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionGenreRepository) FindByGenreID(ctx context.Context, genreID uint) ([]entity.GameRevisionGenre, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("genre_id = ?", genreID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevisionGenre, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionGenreRepository) FindByGameRevisionIDAndGenreIDs(ctx context.Context, gameRevisionID uint, genreIDs []uint) ([]entity.GameRevisionGenre, error) {
	if len(genreIDs) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).
		Where("game_revision_id = ?", gameRevisionID).
		WhereIn("genre_id in (?)", genreIDs).
		Select()

	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevisionGenre, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
