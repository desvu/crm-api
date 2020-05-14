package postgres

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/errors"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type GameRevisionMediaRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) GameRevisionMediaRepository {
	return GameRevisionMediaRepository{
		h: env.Handler,
	}
}

func (r GameRevisionMediaRepository) Create(ctx context.Context, i *entity.GameRevisionMedia) error {
	model := newModel(i)

	_, err := r.h.ModelContext(ctx, model).Insert()
	if err != nil {
		return errors.NewInternal(err)
	}

	*i = *model.Convert()
	return nil
}

func (r GameRevisionMediaRepository) CreateMultiple(ctx context.Context, items []entity.GameRevisionMedia) error {
	if len(items) == 0 {
		return nil
	}

	models := make([]model, len(items))
	for i := range items {
		m := newModel(&items[i])
		models[i] = *m
	}

	_, err := r.h.ModelContext(ctx, &models).Insert()
	if err != nil {
		return errors.NewInternal(err)
	}

	for i := range models {
		items[i] = *models[i].Convert()
	}

	return nil
}

func (r GameRevisionMediaRepository) Delete(ctx context.Context, i *entity.GameRevisionMedia) error {
	model := newModel(i)

	_, err := r.h.ModelContext(ctx, model).WherePK().Delete()
	if err != nil {
		return errors.NewInternal(err)
	}

	*i = *model.Convert()
	return nil
}

func (r GameRevisionMediaRepository) DeleteMultiple(ctx context.Context, items []entity.GameRevisionMedia) error {
	if len(items) == 0 {
		return nil
	}

	models := make([]model, len(items))
	for i := range items {
		m := newModel(&items[i])
		models[i] = *m
	}

	_, err := r.h.ModelContext(ctx, &models).Delete()
	if err != nil {
		return errors.NewInternal(err)
	}

	return nil
}

func (r GameRevisionMediaRepository) FindByRevisionID(ctx context.Context, revisionID uint) ([]entity.GameRevisionMedia, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("revision_id = ?", revisionID).Select()
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	entities := make([]entity.GameRevisionMedia, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionMediaRepository) FindByRevisionIDs(ctx context.Context, revisionIDs []uint) ([]entity.GameRevisionMedia, error) {
	if len(revisionIDs) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).WhereIn("revision_id in (?)", revisionIDs).Select()
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	entities := make([]entity.GameRevisionMedia, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
