package postgres

import (
	"context"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/qilin/crm-api/internal/domain/entity"
	"github.com/qilin/crm-api/internal/domain/errors"
	"github.com/qilin/crm-api/internal/env"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
	"github.com/qilin/crm-api/pkg/transactor"
)

type StorefrontRepository struct {
	h  sql.Handler
	tx *transactor.Transactor
}

func New(env *env.Postgres, tx *transactor.Transactor) StorefrontRepository {
	return StorefrontRepository{
		h:  env.Handler,
		tx: tx,
	}
}

// Create inserts new storefront in db
func (r StorefrontRepository) Create(ctx context.Context, i *entity.Storefront) error {
	sf, err := newStorefront(i)
	if err != nil {
		return err
	}

	err = r.tx.Transact(ctx, func(ctx context.Context) error {
		_, err = r.h.ModelContext(ctx, sf).Insert()
		if err != nil {
			return errors.NewInternal(err)
		}

		sf.Version.StorefrontID = sf.ID

		_, err = r.h.ModelContext(ctx, &sf.Version).Insert()
		if err != nil {
			return errors.NewInternal(err)
		}

		return nil
	})
	if err != nil {
		return err
	}

	*i = *sf.Convert()
	return nil
}

// Update inserts new storefront version in db
func (r StorefrontRepository) Update(ctx context.Context, i *entity.Storefront) error {
	sf, err := newStorefront(i)
	if err != nil {
		return err
	}

	err = r.tx.Transact(ctx, func(ctx context.Context) error {
		_, err = r.h.ModelContext(ctx, sf).WherePK().Update()
		if err != nil {
			return errors.NewInternal(err)
		}

		_, err = r.h.ModelContext(ctx, &sf.Version).Insert()
		if err != nil {
			return errors.NewInternal(err)
		}

		return nil
	})
	if err != nil {
		return err
	}

	*i = *sf.Convert()
	return nil
}

// Delete storefront from db
func (r StorefrontRepository) Delete(ctx context.Context, id uint) error {
	_, err := r.h.ModelContext(ctx, (*storefront)(nil)).Where("id = ?", id).Delete()
	if err != nil {
		return errors.NewInternal(err)
	}
	return nil
}

func (r StorefrontRepository) Activate(ctx context.Context, id, version uint) error {
	act := &activation{
		StorefrontID: id,
		VersionID:    version,
	}

	_, err := r.h.ModelContext(ctx, act).Insert()
	if err != nil {
		return errors.NewInternal(err)
	}

	return nil
}

func (r StorefrontRepository) FindByID(ctx context.Context, id uint) (*entity.Storefront, error) {
	sf := new(storefront)
	err := r.h.ModelContext(ctx, sf).
		Column("sf.*").
		ColumnExpr("((?) = sf.id) as is_active", r.lastActive()).
		Where("sf.id=?", id).
		Relation("Version").Order("version desc").
		First()

	if err == pg.ErrNoRows {
		return nil, errors.StoreFrontNotFound
	}

	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return sf.Convert(), nil
}

func (r StorefrontRepository) FindByIDAndVersion(ctx context.Context, id, version uint) (*entity.Storefront, error) {
	sf := new(storefront)
	err := r.h.ModelContext(ctx, sf).
		Column("sf.*").
		Where("sf.id=?", id).
		Relation("Version", func(q *orm.Query) (*orm.Query, error) {
			return q.JoinOn("id = ?", version), nil
		}).
		First()
	if err == pg.ErrNoRows {
		return nil, errors.StoreFrontNotFound
	}
	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return sf.Convert(), nil
}

func (r StorefrontRepository) FindAll(ctx context.Context) ([]*entity.Storefront, error) {
	var sf []storefront
	err := r.h.ModelContext(ctx, &sf).
		DistinctOn("sf.id").Order("sf.id").
		Column("sf.*").
		ColumnExpr("((?) = sf.id) as is_active", r.lastActive()).
		Relation("Version").Order("version desc").
		Select()

	if err != nil {
		return nil, errors.NewInternal(err)
	}

	result := make([]*entity.Storefront, len(sf))
	for i := range sf {
		result[i] = sf[i].Convert()
	}

	return result, nil
}
func (r StorefrontRepository) lastActive() *orm.Query {
	return r.h.GetConnection().Model((*activation)(nil)).Column("storefront_id").Order("timestamp desc").Limit(1)
}

func (r StorefrontRepository) FindActive(ctx context.Context) (*entity.Storefront, error) {
	sf := new(storefront)
	err := r.h.ModelContext(ctx, sf).
		Column("sf.*").
		ColumnExpr("true as is_active").
		Where("sf.id=(?)", r.lastActive()).
		Relation("Version").Order("version desc").
		First()

	if err == pg.ErrNoRows {
		return nil, errors.StoreFrontNotFound
	}

	if err != nil {
		return nil, errors.NewInternal(err)
	}

	return sf.Convert(), nil
}
