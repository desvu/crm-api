package postgres

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type GameRevisionLocalizationRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) GameRevisionLocalizationRepository {
	return GameRevisionLocalizationRepository{
		h: env.Handler,
	}
}

func (r GameRevisionLocalizationRepository) Create(ctx context.Context, i *entity.Localization) error {
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

func (r GameRevisionLocalizationRepository) CreateMultiple(ctx context.Context, items []entity.Localization) error {
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

func (r GameRevisionLocalizationRepository) Update(ctx context.Context, i *entity.Localization) error {
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

func (r GameRevisionLocalizationRepository) UpdateMultiple(ctx context.Context, items []entity.Localization) error {
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

	_, err := r.h.ModelContext(ctx, &models).Update()
	if err != nil {
		return err
	}

	for i := range models {
		items[i] = *models[i].Convert()
	}

	return nil
}

func (r GameRevisionLocalizationRepository) Delete(ctx context.Context, i *entity.Localization) error {
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

func (r GameRevisionLocalizationRepository) DeleteMultiple(ctx context.Context, items []entity.Localization) error {
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

	_, err := r.h.ModelContext(ctx, &models).Delete()
	if err != nil {
		return err
	}

	return nil
}

func (r GameRevisionLocalizationRepository) FindByID(ctx context.Context, id uint) (*entity.Localization, error) {
	model := new(model)

	err := r.h.ModelContext(ctx, model).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	return model.Convert(), nil
}

func (r GameRevisionLocalizationRepository) FindByIDs(ctx context.Context, ids []uint) ([]entity.Localization, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).WhereIn("id in (?)", ids).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Localization, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionLocalizationRepository) FindByGameRevisionID(ctx context.Context, gameRevisionID uint) ([]entity.Localization, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).Where("game_revision_id = ?", gameRevisionID).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Localization, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionLocalizationRepository) FindByGameRevisionIDs(ctx context.Context, gameRevisionIDs []uint) ([]entity.Localization, error) {
	if len(gameRevisionIDs) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).WhereIn("game_revision_id in (?)", gameRevisionIDs).Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Localization, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionLocalizationRepository) FindByGameRevisionIDAndLanguage(ctx context.Context, gameRevisionID uint, langs []string) ([]entity.Localization, error) {
	if len(langs) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).
		Where("game_revision_id = ?", gameRevisionID).
		WhereIn("language in (?)", langs).
		Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.Localization, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}
