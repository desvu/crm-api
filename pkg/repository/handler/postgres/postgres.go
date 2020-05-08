package postgres

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
	"github.com/qilin/crm-api/pkg/transactor"
)

type Handler struct {
	conn             *pg.DB
	transactionStore *transactor.Store
}

func (h *Handler) ModelContext(c context.Context, model ...interface{}) *orm.Query {
	th, ok, err := h.transactionStore.GetHandler(c, h)
	if ok {
		if err != nil {
			return nil
		}

		return th.ModelContext(c, model...)
	}

	return h.conn.ModelContext(c, model...)
}

func (h *Handler) ExecContext(c context.Context, query interface{}, params ...interface{}) (orm.Result, error) {
	th, ok, err := h.transactionStore.GetHandler(c, h)
	if ok {
		if err != nil {
			return nil, err
		}

		return th.ExecContext(c, query, params...)
	}

	return h.conn.ExecContext(c, query, params...)
}

func (h *Handler) ExecOneContext(c context.Context, query interface{}, params ...interface{}) (orm.Result, error) {
	th, ok, err := h.transactionStore.GetHandler(c, h)
	if ok {
		if err != nil {
			return nil, err
		}

		return th.ExecOneContext(c, query, params...)
	}

	return h.conn.ExecOneContext(c, query, params...)
}

func (h *Handler) QueryContext(c context.Context, model, query interface{}, params ...interface{}) (orm.Result, error) {
	th, ok, err := h.transactionStore.GetHandler(c, h)
	if ok {
		if err != nil {
			return nil, err
		}

		return th.QueryContext(c, model, query, params...)
	}

	return h.conn.QueryContext(c, model, query, params...)
}

func (h *Handler) QueryOneContext(c context.Context, model, query interface{}, params ...interface{}) (orm.Result, error) {
	th, ok, err := h.transactionStore.GetHandler(c, h)
	if ok {
		if err != nil {
			return nil, err
		}

		return th.QueryOneContext(c, model, query, params...)
	}

	return h.conn.QueryOneContext(c, model, query, params...)
}

func (h *Handler) Begin() (sql.TransactionHandler, error) {
	tx, err := h.conn.Begin()
	if err != nil {
		return nil, err
	}

	return &TransactionHandler{
		tx: tx,
	}, nil
}

func (h *Handler) GetConnection() *pg.DB {
	return h.conn
}

type TransactionHandler struct {
	tx *pg.Tx
}

func (h *TransactionHandler) ModelContext(c context.Context, model ...interface{}) *orm.Query {
	return h.tx.ModelContext(c, model...)
}

func (h *TransactionHandler) ExecContext(c context.Context, query interface{}, params ...interface{}) (orm.Result, error) {
	return h.tx.ExecContext(c, query, params...)
}

func (h *TransactionHandler) ExecOneContext(c context.Context, query interface{}, params ...interface{}) (orm.Result, error) {
	return h.tx.ExecOneContext(c, query, params...)
}

func (h *TransactionHandler) QueryContext(c context.Context, model, query interface{}, params ...interface{}) (orm.Result, error) {
	return h.tx.QueryContext(c, model, query, params...)
}

func (h *TransactionHandler) QueryOneContext(c context.Context, model, query interface{}, params ...interface{}) (orm.Result, error) {
	return h.tx.QueryOneContext(c, model, query, params...)
}

func (h *TransactionHandler) Commit() error {
	return h.tx.Commit()
}

func (h *TransactionHandler) Rollback() error {
	return h.tx.Rollback()
}

func (h *TransactionHandler) GetConnection() *pg.DB {
	panic("transaction handler dont returned basic connection")
}

type Config struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

func New(cfg Config, transactionStore *transactor.Store) *Handler {
	return NewOptions(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Database: cfg.Database,
		User:     cfg.User,
		Password: cfg.Password,
	}, transactionStore)
}

func NewOptions(opts *pg.Options, transactionStore *transactor.Store) *Handler {
	return &Handler{
		transactionStore: transactionStore,
		conn:             pg.Connect(opts),
	}
}
