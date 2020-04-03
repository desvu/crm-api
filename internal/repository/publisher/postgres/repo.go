package postgres

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
)

type PublisherRepository struct {
	db *pg.DB
}

func New(env *env.Postgres) PublisherRepository {
	return PublisherRepository{db: env.DB}
}

func (r PublisherRepository) Create(ctx context.Context, i *entity.Publisher) error {
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

func (r PublisherRepository) Update(ctx context.Context, i *entity.Publisher) error {
	model, err := newModel(i)
	if err != nil {
		return err
	}

	_, err = r.db.ModelContext(ctx, model).WherePK().Update()
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

	_, err = r.db.ModelContext(ctx, model).WherePK().Delete()
	if err != nil {
		return err
	}

	*i = *model.Convert()
	return nil
}

func (r PublisherRepository) FindByID(ctx context.Context, id uint) (*entity.Publisher, error) {
	model := new(model)

	err := r.db.ModelContext(ctx, model).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return model.Convert(), nil
}

func (r PublisherRepository) FindByIDs(ctx context.Context, ids []uint) ([]entity.Publisher, error) {
	var models []model

	err := r.db.ModelContext(ctx, &models).Where("id in (?)", pg.In(ids)).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Publisher, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
