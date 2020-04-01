package postgres

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/repository"
	"github.com/qilin/crm-api/internal/env"
)

type GameRepository struct {
	conn *pg.DB
}

func New(env *env.Postgres) repository.IGameRepository {
	return GameRepository{conn: env.Conn}
}

func (r GameRepository) Create(ctx context.Context, i *entity.Game) error {
	m, err := newModel(i)
	if err != nil {
		return err
	}

	if err = r.conn.Insert(m); err != nil {
		return err
	}

	*i = *m.Convert()
	return nil
}

func (r GameRepository) FindByID(ctx context.Context, id uint) (*entity.Game, error) {
	user := model{ID: id}
	if err := r.conn.Select(&user); err != nil {
		return nil, err
	}

	return user.Convert(), nil
}
