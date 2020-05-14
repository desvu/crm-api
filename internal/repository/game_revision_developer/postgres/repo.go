package postgres

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type GameRevisionDeveloperRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) GameRevisionDeveloperRepository {
	return GameRevisionDeveloperRepository{
		h: env.Handler,
	}
}

func (r GameRevisionDeveloperRepository) Create(ctx context.Context, i *entity.GameRevisionDeveloper) error {
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

func (r GameRevisionDeveloperRepository) CreateMultiple(ctx context.Context, items []entity.GameRevisionDeveloper) error {
	if len(items) == 0 {
		return nil
	}

	models := make([]model, len(items))
	for i := range items {
		m, err := newModel(&items[i])
		if err != nil {
			return err
		}

		models[i] = *m
	}

	_, err := r.h.ModelContext(ctx, &models).Insert()
	if err != nil {
		return err
	}

	for i := range models {
		items[i] = *models[i].Convert()
	}

	return nil
}

func (r GameRevisionDeveloperRepository) Delete(ctx context.Context, i *entity.GameRevisionDeveloper) error {
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

func (r GameRevisionDeveloperRepository) DeleteMultiple(ctx context.Context, items []entity.GameRevisionDeveloper) error {
	if len(items) == 0 {
		return nil
	}

	models := make([]model, len(items))
	for i := range items {
		m, err := newModel(&items[i])
		if err != nil {
			return err
		}

		models[i] = *m
	}

	_, err := r.h.ModelContext(ctx, &models).WherePK().Delete()
	if err != nil {
		return err
	}

	return nil
}

func (r GameRevisionDeveloperRepository) FindByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.GameRevisionDeveloper, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("game_revision_id = ?", gameRevisionID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevisionDeveloper, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionDeveloperRepository) FindByGameRevisionIDs(ctx context.Context, gameRevisionIDs []uint) ([]entity.GameRevisionDeveloper, error) {
	if len(gameRevisionIDs) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).WhereIn("game_revision_id in (?)", gameRevisionIDs).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevisionDeveloper, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionDeveloperRepository) FindByDeveloperID(ctx context.Context, developerID uint) ([]entity.GameRevisionDeveloper, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("developer_id = ?", developerID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevisionDeveloper, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionDeveloperRepository) FindByGameRevisionIDAndDeveloperIDs(ctx context.Context, gameRevisionID uint, developerIDs []uint) ([]entity.GameRevisionDeveloper, error) {
	if len(developerIDs) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).
		Where("game_revision_id = ?", gameRevisionID).
		WhereIn("developer_id in (?)", developerIDs).
		Select()

	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevisionDeveloper, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
