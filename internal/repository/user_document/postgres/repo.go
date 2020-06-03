package postgres

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type UserDocumentRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) UserDocumentRepository {
	return UserDocumentRepository{
		h: env.Handler,
	}
}

func (r UserDocumentRepository) Create(ctx context.Context, i *entity.UserDocument) error {
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

func (r UserDocumentRepository) FindByUserID(ctx context.Context, filter repository.FindUserDocumentsByUserIdData) ([]entity.UserDocument, error) {
	var models []model

	err := r.h.ModelContext(ctx, &models).
		Where("user_id = ?", filter.UserID).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Select()
	if err != nil {
		return nil, err
	}

	entities := make([]entity.UserDocument, len(models))
	for i := range models {
		entities[i] = *models[i].Convert()
	}

	return entities, nil
}

func (r UserDocumentRepository) FindByUserAndDocumentID(ctx context.Context, userId, docId uint) (*entity.UserDocument, error) {
	var model model

	err := r.h.ModelContext(ctx, &model).
		Where("user_id = ?", userId).
		Where("document_id = ?", docId).
		Select()

	if err == pg.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return model.Convert(), nil
}
