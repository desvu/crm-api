package postgres

import (
	"context"

	"github.com/pkg/errors"

	"github.com/go-pg/pg/v9"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type GameStorePublishRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) GameStorePublishRepository {
	return GameStorePublishRepository{
		h: env.Handler,
	}
}

func (r GameStorePublishRepository) Create(ctx context.Context, i *entity.GameStorePublish) error {
	model := newModel(i)

	_, err := r.h.ModelContext(ctx, model).Insert()
	if err != nil {
		return errors.WithStack(err)
	}

	*i = *model.Convert()
	return nil
}

func (r GameStorePublishRepository) Update(ctx context.Context, i *entity.GameStorePublish) error {
	model := newModel(i)

	_, err := r.h.ModelContext(ctx, model).WherePK().Update()
	if err != nil {
		return errors.WithStack(err)
	}

	*i = *model.Convert()
	return nil
}

func (r GameStorePublishRepository) Delete(ctx context.Context, i *entity.GameStorePublish) error {
	model := newModel(i)

	_, err := r.h.ModelContext(ctx, model).WherePK().Delete()
	if err != nil {
		return err
	}

	*i = *model.Convert()
	return nil
}

func (r GameStorePublishRepository) FindByID(ctx context.Context, id uint) (*entity.GameStorePublish, error) {
	model := new(model)

	err := r.h.ModelContext(ctx, model).Where("id = ?", id).Select()

	if err == pg.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return model.Convert(), nil
}

func (r GameStorePublishRepository) FindByGameID(ctx context.Context, gameID string) ([]entity.GameStorePublish, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("game_id = ?", gameID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameStorePublish, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
