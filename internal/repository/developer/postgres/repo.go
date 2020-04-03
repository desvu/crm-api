package postgres

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
)

type DeveloperRepository struct {
	db *pg.DB
}

func New(env *env.Postgres) DeveloperRepository {
	return DeveloperRepository{db: env.DB}
}

func (r DeveloperRepository) Create(ctx context.Context, i *entity.Developer) error {
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

func (r DeveloperRepository) Update(ctx context.Context, i *entity.Developer) error {
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

func (r DeveloperRepository) Delete(ctx context.Context, i *entity.Developer) error {
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

func (r DeveloperRepository) FindByID(ctx context.Context, id uint) (*entity.Developer, error) {
	model := new(model)

	err := r.db.ModelContext(ctx, model).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return model.Convert(), nil
}

func (r DeveloperRepository) FindByIDs(ctx context.Context, ids []uint) ([]entity.Developer, error) {
	var models []model

	err := r.db.ModelContext(ctx, &models).Where("id in (?)", pg.In(ids)).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Developer, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
