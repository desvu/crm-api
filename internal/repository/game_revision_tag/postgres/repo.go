package postgres

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type GameRevisionTagRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) GameRevisionTagRepository {
	return GameRevisionTagRepository{
		h: env.Handler,
	}
}

func (r GameRevisionTagRepository) Create(ctx context.Context, i *entity.GameRevisionTag) error {
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

func (r GameRevisionTagRepository) CreateMultiple(ctx context.Context, items []entity.GameRevisionTag) error {
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
		return err
	}

	for i := range models {
		items[i] = *models[i].Convert()
	}

	return nil
}

func (r GameRevisionTagRepository) Delete(ctx context.Context, i *entity.GameRevisionTag) error {
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

func (r GameRevisionTagRepository) DeleteMultiple(ctx context.Context, items []entity.GameRevisionTag) error {
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
		return err
	}

	return nil
}

func (r GameRevisionTagRepository) FindByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.GameRevisionTag, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("game_revision_id = ?", gameRevisionID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevisionTag, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionTagRepository) FindByGameRevisionIDs(ctx context.Context, gameRevisionIDs []uint) ([]entity.GameRevisionTag, error) {
	if len(gameRevisionIDs) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).WhereIn("game_revision_id in (?)", gameRevisionIDs).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevisionTag, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionTagRepository) FindByTagID(ctx context.Context, tagID uint) ([]entity.GameRevisionTag, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("tag_id = ?", tagID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevisionTag, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionTagRepository) FindByGameRevisionIDAndTagIDs(ctx context.Context, gameRevisionID uint, tagIDs []uint) ([]entity.GameRevisionTag, error) {
	if len(tagIDs) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).
		Where("game_revision_id = ?", gameRevisionID).
		Where("tag_id in (?)", tagIDs).
		Select()

	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevisionTag, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
