package postgres

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/errors"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type DocumentRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) DocumentRepository {
	return DocumentRepository{
		h: env.Handler,
	}
}

func (r DocumentRepository) Create(ctx context.Context, i *entity.Document) error {
	model, err := newModel(i)
	if err != nil {
		return err
	}

	_, err = r.h.ModelContext(ctx, model).Insert()
	if err != nil {
		return errors.NewInternal(err)
	}

	*i = *model.Convert()
	return nil
}

func (r DocumentRepository) Update(ctx context.Context, i *entity.Document) error {
	model, err := newModel(i)
	if err != nil {
		return err
	}

	_, err = r.h.ModelContext(ctx, model).WherePK().Update()
	if err != nil {
		return errors.NewInternal(err)
	}

	*i = *model.Convert()
	return nil
}

func (r DocumentRepository) Delete(ctx context.Context, i *entity.Document) error {
	model, err := newModel(i)
	if err != nil {
		return err
	}

	_, err = r.h.ModelContext(ctx, model).WherePK().Delete()
	if err != nil {
		return errors.NewInternal(err)
	}

	*i = *model.Convert()
	return nil
}

func (r DocumentRepository) FindByID(ctx context.Context, id uint) (*entity.Document, error) {
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

func (r DocumentRepository) FindByIDs(ctx context.Context, ids []uint) ([]entity.Document, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).WhereIn("id in (?)", ids).Select()
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	entities := make([]entity.Document, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r DocumentRepository) FindByFilter(ctx context.Context, filter *repository.FindByFilterDocumentData) ([]entity.Document, error) {
	if filter.Limit == 0 {
		return nil, nil
	}

	var models []model

	q := r.h.ModelContext(ctx, &models)

	if filter.OnlyActivated {
		// todo
	}

	err := q.Limit(filter.Limit).Offset(filter.Offset).Select()

	if err != nil {
		return nil, errors.NewInternal(err)
	}

	entities := make([]entity.Document, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r DocumentRepository) CountByFilter(ctx context.Context, filter *repository.FindByFilterDocumentData) (int, error) {
	var models []model
	q := r.h.ModelContext(ctx, &models)

	if filter.OnlyActivated {
		q.Where("activated_at NOT NULL")
	}

	count, err := q.Count()

	if err != nil {
		return 0, errors.NewInternal(err)
	}

	return count, nil
}
