package postgres

import (
	"context"

	"github.com/pkg/errors"

	"github.com/qilin/crm-api/internal/domain/enum/game_revision"

	"github.com/go-pg/pg/v9"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type GameRevisionRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) GameRevisionRepository {
	return GameRevisionRepository{
		h: env.Handler,
	}
}

func (r GameRevisionRepository) Create(ctx context.Context, i *entity.GameRevision) error {
	model, err := newModel(i)
	if err != nil {
		return err
	}

	_, err = r.h.ModelContext(ctx, model).Insert()
	if err != nil {
		return errors.WithStack(err)
	}

	*i = *model.Convert()
	return nil
}

func (r GameRevisionRepository) Update(ctx context.Context, i *entity.GameRevision) error {
	model, err := newModel(i)
	if err != nil {
		return err
	}

	_, err = r.h.ModelContext(ctx, model).WherePK().Update()
	if err != nil {
		return errors.WithStack(err)
	}

	*i = *model.Convert()
	return nil
}

func (r GameRevisionRepository) Delete(ctx context.Context, i *entity.GameRevision) error {
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

func (r GameRevisionRepository) FindByID(ctx context.Context, id uint) (*entity.GameRevision, error) {
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

func (r GameRevisionRepository) FindByIDs(ctx context.Context, ids []uint) ([]entity.GameRevision, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("id in (?)", pg.In(ids)).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevision, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionRepository) FindByGameID(ctx context.Context, gameID string) ([]entity.GameRevision, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("game_id = ?", gameID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevision, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionRepository) FindLastPublishedByGameID(ctx context.Context, gameID string) (*entity.GameRevision, error) {
	model := new(model)

	err := r.h.ModelContext(ctx, model).
		Where("game_id = ?", gameID).
		Where("status = ?", game_revision.StatusPublished.Value()).
		Order("published_at desc").
		Select()

	if err == pg.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return model.Convert(), nil
}

func (r GameRevisionRepository) FindDraftByGameID(ctx context.Context, gameID string) (*entity.GameRevision, error) {
	model := new(model)

	err := r.h.ModelContext(ctx, model).
		Where("game_id = ?", gameID).
		Where("status = ?", game_revision.StatusDraft.Value()).
		Select()

	if err == pg.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return model.Convert(), nil
}
