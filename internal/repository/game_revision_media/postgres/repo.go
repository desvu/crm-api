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

func (r GameRevisionMediaRepository) Delete(ctx context.Context, i *entity.GameRevisionMedia) error {
	model := newModel(i)

	_, err := r.h.ModelContext(ctx, model).WherePK().Delete()
	if err != nil {
		return errors.NewInternal(err)
	}

	*i = *model.Convert()
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
