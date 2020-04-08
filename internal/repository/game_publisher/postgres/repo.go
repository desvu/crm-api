package postgres

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type GamePublisherRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) GamePublisherRepository {
	return GamePublisherRepository{
		h: env.Handler,
	}
}

func (r GamePublisherRepository) Create(ctx context.Context, i *entity.GamePublisher) error {
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

func (r GamePublisherRepository) CreateMultiple(ctx context.Context, items []entity.GamePublisher) error {
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

func (r GamePublisherRepository) Delete(ctx context.Context, i *entity.GamePublisher) error {
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

func (r GamePublisherRepository) DeleteMultiple(ctx context.Context, items []entity.GamePublisher) error {
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

func (r GamePublisherRepository) FindByGameID(ctx context.Context, gameID uint) ([]entity.GamePublisher, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("game_id = ?", gameID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GamePublisher, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GamePublisherRepository) FindByPublisherID(ctx context.Context, publisherID uint) ([]entity.GamePublisher, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("publisher_id = ?", publisherID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.GamePublisher, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GamePublisherRepository) FindByGameIDAndPublisherIDs(ctx context.Context, gameID uint, publisherIDs []uint) ([]entity.GamePublisher, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).
		Where("game_id = ?", gameID).
		Where("publisher_id in (?)", publisherIDs).
		Select()

	if err != nil {
		return nil, err
	}

	entities := make([]entity.GamePublisher, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
