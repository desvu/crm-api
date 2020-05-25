package postgres

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type GameRevisionRatingRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) GameRevisionRatingRepository {
	return GameRevisionRatingRepository{
		h: env.Handler,
	}
}

func (r GameRevisionRatingRepository) Create(ctx context.Context, i *entity.Rating) error {
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

func (r GameRevisionRatingRepository) CreateMultiple(ctx context.Context, items []entity.Rating) error {
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

func (r GameRevisionRatingRepository) Update(ctx context.Context, i *entity.Rating) error {
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

func (r GameRevisionRatingRepository) UpdateMultiple(ctx context.Context, items []entity.Rating) error {
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

func (r GameRevisionRatingRepository) Delete(ctx context.Context, i *entity.Rating) error {
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

func (r GameRevisionRatingRepository) DeleteMultiple(ctx context.Context, items []entity.Rating) error {
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

func (r GameRevisionRatingRepository) FindByID(ctx context.Context, id uint) (*entity.Rating, error) {
	model := new(model)

	err := r.h.ModelContext(ctx, model).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return model.Convert(), nil
}

func (r GameRevisionRatingRepository) FindByIDs(ctx context.Context, ids []uint) ([]entity.Rating, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).WhereIn("id in (?)", ids).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Rating, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionRatingRepository) FindByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.Rating, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("game_revision_id = ?", gameRevisionID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Rating, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionRatingRepository) FindByGameRevisionIDAndAgency(ctx context.Context, gameRevisionID uint, agencies []uint8) ([]entity.Rating, error) {
	if len(agencies) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).
		Where("game_revision_id = ?", gameRevisionID).
		WhereIn("agency in (?)", agencies).
		Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Rating, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
