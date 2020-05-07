package postgres

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/errors"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type GameMediaRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) GameMediaRepository {
	return GameMediaRepository{
		h: env.Handler,
	}
}

func (r GameMediaRepository) Create(ctx context.Context, i *entity.GameMedia) error {
	model := newModel(i)

	_, err := r.h.ModelContext(ctx, model).Insert()
	if err != nil {
		return errors.NewInternal(err)
	}

	*i = *model.Convert()
	return nil
}

func (r GameMediaRepository) Update(ctx context.Context, i *entity.GameMedia) error {
	model := newModel(i)

	_, err := r.h.ModelContext(ctx, model).WherePK().Update()
	if err != nil {
		return errors.NewInternal(err)
	}

	*i = *model.Convert()
	return nil
}

func (r GameMediaRepository) Delete(ctx context.Context, i *entity.GameMedia) error {
	model := newModel(i)

	_, err := r.h.ModelContext(ctx, model).WherePK().Delete()
	if err != nil {
		return errors.NewInternal(err)
	}

	*i = *model.Convert()
	return nil
}

func (r GameMediaRepository) FindByID(ctx context.Context, id uint) (*entity.GameMedia, error) {
	model := new(model)

	err := r.h.ModelContext(ctx, model).Where("id = ?", id).Select()

	if err == pg.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return model.Convert(), nil
}

func (r GameMediaRepository) FindByIDs(ctx context.Context, ids []uint) ([]entity.GameMedia, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("id in (?)", pg.In(ids)).Select()
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	entities := make([]entity.GameMedia, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
