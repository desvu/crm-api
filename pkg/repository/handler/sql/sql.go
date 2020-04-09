package sql

import (
	"context"

	"github.com/go-pg/pg/v9"

	"github.com/go-pg/pg/v9/orm"
)

type Handler interface {
	ModelContext(c context.Context, model ...interface{}) *orm.Query
	ExecContext(c context.Context, query interface{}, params ...interface{}) (orm.Result, error)
	ExecOneContext(c context.Context, query interface{}, params ...interface{}) (orm.Result, error)
	QueryContext(c context.Context, model, query interface{}, params ...interface{}) (orm.Result, error)
	QueryOneContext(c context.Context, model, query interface{}, params ...interface{}) (orm.Result, error)

	GetConnection() *pg.DB
}

type TransactionBeginner interface {
	Begin() (TransactionHandler, error)
}

type TransactionHandler interface {
	Handler
	Commit() error
	Rollback() error
}
