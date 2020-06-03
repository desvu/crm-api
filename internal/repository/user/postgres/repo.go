package postgres

import (
	"context"

	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type UserRepository struct {
	h sql.Handler
}

func New(env *env.Postgres) UserRepository {
	return UserRepository{
		h: env.Handler,
	}
}

func (r UserRepository) FindByExternalID(ctx context.Context, externalID string) (*entity.User, error) {
	model := new(model)

	err := r.h.ModelContext(ctx, model).Where("external_id = ?", externalID).Select()
	if err != nil {
		return nil, err
	}

	return model.Convert(), nil
}
