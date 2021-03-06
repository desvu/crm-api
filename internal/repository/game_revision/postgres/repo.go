package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/enum/enum"
	"github.com/qilin/crm-api/internal/domain/enum/game_revision"
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/errors"
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
		return errors.NewInternal(err)
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
		return errors.NewInternal(err)
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
		return errors.NewInternal(err)
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
		return nil, errors.NewInternal(err)
	}

	return model.Convert(), nil
}

func (r GameRevisionRepository) FindByIDs(ctx context.Context, ids []uint) ([]entity.GameRevision, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).WhereIn("id in (?)", ids).Select()
	if err != nil {
		return nil, errors.NewInternal(err)
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
		return nil, errors.NewInternal(err)
	}

	entities := make([]entity.GameRevision, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionRepository) FindLastByGameIDs(ctx context.Context, gameIDs []string) ([]entity.GameRevision, error) {
	if len(gameIDs) == 0 {
		return nil, nil
	}

	var models []model

	err := r.h.ModelContext(ctx, &models).
		DistinctOn("game_id").
		WhereIn("game_id in (?)", gameIDs).
		Order("game_id", "id desc").
		Select()

	if err != nil {
		return nil, errors.NewInternal(err)
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
		First()

	if err == pg.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return model.Convert(), nil
}

func (r GameRevisionRepository) FindDraftByGameID(ctx context.Context, gameID string) (*entity.GameRevision, error) {
	model := new(model)

	err := r.h.ModelContext(ctx, model).
		Where("game_id = ?", gameID).
		Where("status = ?", game_revision.StatusDraft.Value()).
		First()

	if err == pg.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return model.Convert(), nil
}

func (r GameRevisionRepository) FindByIDAndGameID(ctx context.Context, id uint, gameID string) (*entity.GameRevision, error) {
	model := new(model)

	err := r.h.ModelContext(ctx, model).
		Where("id = ?", id).
		Where("game_id = ?", gameID).
		Select()

	if err == pg.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return model.Convert(), nil
}

func (r GameRevisionRepository) FindPublishedByGameIDs(ctx context.Context, gameIDs []string) ([]string, error) {
	if len(gameIDs) == 0 {
		return nil, nil
	}

	var res []string
	err := r.h.ModelContext(ctx, (*model)(nil)).
		Column("game_id").
		WhereIn("game_id in (?)", gameIDs).
		Where("status = ?", game_revision.StatusPublished.Value()).
		Group("game_id").
		Select(&res)

	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return res, nil
}

func (r GameRevisionRepository) FindByFilter(ctx context.Context, filter *repository.FindByFilterGameRevisionData) ([]entity.GameRevision, error) {
	if filter.Limit == 0 {
		return nil, nil
	}

	var models []model

	q := r.h.ModelContext(ctx, &models).
		ColumnExpr("model.*").
		DistinctOn("model.game_id").
		Join("join games on games.id = model.game_id")

	if len(filter.Title) != 0 {
		q.Where("games.title ilike ?", "%"+filter.Title+"%")
	}

	if filter.OnlyPublished {
		q.Where("model.status = ?", game_revision.StatusPublished.Value())
	}

	if len(filter.GenreIDs) > 0 {
		q.Join("join game_revision_genres on game_revision_genres.game_revision_id = model.id").
			WhereIn("game_revision_genres.genre_id in (?)", filter.GenreIDs)
	}

	if len(filter.FeatureIDs) > 0 {
		q.Join("join game_revision_features on game_revision_features.game_revision_id = model.id").
			WhereIn("game_revision_features.feature_id in (?)", filter.FeatureIDs)
	}

	orderType := "desc"
	if filter.OrderType == enum.SortOrderAsc {
		orderType = "asc"
	}

	switch filter.OrderBy {
	case enum.SortOrderColumnReleaseDate:
		q.Order("model.game_id", fmt.Sprintf("%s %s", "model.release_date", orderType))
	case enum.SortOrderColumnName:
		q.Order("model.game_id", fmt.Sprintf("%s %s", "games.title", orderType))
	default:
		q.Order("model.game_id", fmt.Sprintf("%s %s", "model.id", orderType))
	}

	err := q.Limit(filter.Limit).Offset(filter.Offset).Select()

	if err != nil {
		return nil, errors.NewInternal(err)
	}

	entities := make([]entity.GameRevision, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r GameRevisionRepository) CountByFilter(ctx context.Context, filter *repository.FindByFilterGameRevisionData) (int, error) {
	var models []model

	q := r.h.ModelContext(ctx, &models).
		ColumnExpr("model.*").
		DistinctOn("model.game_id").
		Join("join games on games.id = model.game_id")

	if len(filter.Title) != 0 {
		q.Where("games.title ilike ?", "%"+filter.Title+"%")
	}

	if filter.OnlyPublished {
		q.Where("model.status = ?", game_revision.StatusPublished.Value())
	}

	if len(filter.GenreIDs) > 0 {
		q.Join("join game_revision_genres on game_revision_genres.game_revision_id = model.id").
			WhereIn("game_revision_genres.genre_id in (?)", filter.GenreIDs)
	}

	if len(filter.FeatureIDs) > 0 {
		q.Join("join game_revision_features on game_revision_features.game_revision_id = model.id").
			WhereIn("game_revision_features.feature_id in (?)", filter.FeatureIDs)
	}

	count, err := q.Count()

	if err != nil {
		return 0, errors.NewInternal(err)
	}

	return count, nil
}
