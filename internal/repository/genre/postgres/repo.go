package postgres

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type GenreRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) GenreRepository {
	return GenreRepository{
		h: env.Handler,
	}
}

func (r GenreRepository) Create(ctx context.Context, i *entity.Genre) error {
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

func (r GenreRepository) Update(ctx context.Context, i *entity.Genre) error {
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

func (r GenreRepository) Delete(ctx context.Context, i *entity.Genre) error {
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

func (r GenreRepository) FindAll(ctx context.Context) ([]entity.Genre, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Select()
	if err != nil {
		return nil, err
	}

	result := make([]entity.Genre, len(models))
	for i := range models {
		result[i] = *models[i].Convert()
	}

	return result, nil
}

func (r GenreRepository) FindByID(ctx context.Context, id uint) (*entity.Genre, error) {
	model := new(model)

	err := r.h.ModelContext(ctx, model).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return model.Convert(), nil
}

func (r GenreRepository) FindByIDs(ctx context.Context, ids []uint) ([]entity.Genre, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).WhereIn("id in (?)", ids).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Genre, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
