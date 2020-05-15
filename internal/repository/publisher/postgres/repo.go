package postgres

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/errors"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type PublisherRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) PublisherRepository {
	return PublisherRepository{
		h: env.Handler,
	}
}

func (r PublisherRepository) Create(ctx context.Context, i *entity.Publisher) error {
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

func (r PublisherRepository) Update(ctx context.Context, i *entity.Publisher) error {
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

func (r PublisherRepository) Delete(ctx context.Context, i *entity.Publisher) error {
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

func (r PublisherRepository) FindByID(ctx context.Context, id uint) (*entity.Publisher, error) {
	model := new(model)

	err := r.h.ModelContext(ctx, model).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return model.Convert(), nil
}

func (r PublisherRepository) FindByIDs(ctx context.Context, ids []uint) ([]entity.Publisher, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).WhereIn("id in (?)", ids).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Publisher, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r PublisherRepository) FindByFilter(ctx context.Context, data *repository.FindByFilterPublisherData) ([]entity.Publisher, error) {
	var models []model

	q := r.h.ModelContext(ctx, &models).
		Limit(data.Limit)

	if data.Offset != 0 {
		q.Offset(data.Offset)
	}

	if err := q.Select(); err != nil {
		return nil, errors.NewInternal(err)
	}

	entities := make([]entity.Publisher, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
