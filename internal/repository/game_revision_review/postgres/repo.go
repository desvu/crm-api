package postgres

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type GameRevisionReviewRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) GameRevisionReviewRepository {
	return GameRevisionReviewRepository{
		h: env.Handler,
	}
}

func (r GameRevisionReviewRepository) Create(ctx context.Context, i *entity.Review) error {
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

func (r GameRevisionReviewRepository) CreateMultiple(ctx context.Context, items []entity.Review) error {
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

func (r GameRevisionReviewRepository) Update(ctx context.Context, i *entity.Review) error {
	model, err := newModel(i)
	if err != nil {
		return err
	}

	_, err = r.h.ModelContext(ctx, model).WherePK().Update()
	if err != nil {
		return err
	}

	*i = *model.Convert()
	return nil
}

func (r GameRevisionReviewRepository) UpdateMultiple(ctx context.Context, items []entity.Review) error {
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

	_, err := r.h.ModelContext(ctx, &models).Update()
	if err != nil {
		return err
	}

	for i := range models {
		items[i] = *models[i].Convert()
	}

	return nil
}

func (r GameRevisionReviewRepository) Delete(ctx context.Context, i *entity.Review) error {
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

func (r GameRevisionReviewRepository) DeleteMultiple(ctx context.Context, items []entity.Review) error {
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

func (r GameRevisionReviewRepository) FindByID(ctx context.Context, id uint) (*entity.Review, error) {
	model := new(model)

	err := r.h.ModelContext(ctx, model).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return model.Convert(), nil
}

func (r GameRevisionReviewRepository) FindByIDs(ctx context.Context, ids []uint) ([]entity.Review, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).WhereIn("id in (?)", ids).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Review, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionReviewRepository) FindByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.Review, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("game_revision_id = ?", gameRevisionID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Review, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
