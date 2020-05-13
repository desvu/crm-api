package postgres

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type GameRevisionPublisherRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) GameRevisionPublisherRepository {
	return GameRevisionPublisherRepository{
		h: env.Handler,
	}
}

func (r GameRevisionPublisherRepository) Create(ctx context.Context, i *entity.GameRevisionPublisher) error {
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

func (r GameRevisionPublisherRepository) CreateMultiple(ctx context.Context, items []entity.GameRevisionPublisher) error {
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

	_, err := r.h.ModelContext(ctx, models).Insert()
	if err != nil {
		return err
	}

	for i := range models {
		items[i] = *models[i].Convert()
	}

	return nil
}

func (r GameRevisionPublisherRepository) Delete(ctx context.Context, i *entity.GameRevisionPublisher) error {
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

func (r GameRevisionPublisherRepository) DeleteMultiple(ctx context.Context, items []entity.GameRevisionPublisher) error {
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

	_, err := r.h.ModelContext(ctx, models).Delete()
	if err != nil {
		return err
	}

	return nil
}

func (r GameRevisionPublisherRepository) FindByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.GameRevisionPublisher, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("game_revision_id = ?", gameRevisionID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevisionPublisher, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionPublisherRepository) FindByGameRevisionIDs(ctx context.Context, gameRevisionIDs []uint) ([]entity.GameRevisionPublisher, error) {
	if len(gameRevisionIDs) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).WhereIn("game_revision_id in (?)", gameRevisionIDs).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevisionPublisher, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionPublisherRepository) FindByPublisherID(ctx context.Context, publisherID uint) ([]entity.GameRevisionPublisher, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("publisher_id = ?", publisherID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevisionPublisher, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionPublisherRepository) FindByGameRevisionIDAndPublisherIDs(ctx context.Context, gameRevisionID uint, publisherIDs []uint) ([]entity.GameRevisionPublisher, error) {
	if len(publisherIDs) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).
		Where("game_revision_id = ?", gameRevisionID).
		WhereIn("publisher_id in (?)", publisherIDs).
		Select()

	if err != nil {
		return nil, err
	}

	entities := make([]entity.GameRevisionPublisher, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
