package postgres

import (
	"context"

	"github.com/qilin/crm-api/pkg/repository/handler/sql"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
)

type FeatureRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) FeatureRepository {
	return FeatureRepository{
		h: env.Handler,
	}
}

func (r FeatureRepository) Create(ctx context.Context, i *entity.Feature) error {
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

func (r FeatureRepository) Update(ctx context.Context, i *entity.Feature) error {
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

func (r FeatureRepository) Delete(ctx context.Context, i *entity.Feature) error {
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

func (r FeatureRepository) FindByID(ctx context.Context, id uint) (*entity.Feature, error) {
	model := new(model)

	err := r.h.ModelContext(ctx, model).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return model.Convert(), nil
}

func (r FeatureRepository) FindByIDs(ctx context.Context, ids []uint) ([]entity.Feature, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).WhereIn("id in (?)", ids).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Feature, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
