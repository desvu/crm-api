package transaction

import (
	"context"
	"errors"
	"sync"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/qilin/crm-api/pkg/context/transact"
)

var ErrTransactionSQLHandlerNotFound = errors.New("transaction sql handler not found")

type Store struct {
	handlers map[string]*pg.Tx
	mx       sync.RWMutex
}

func New() *Store {
	return &Store{
		handlers: make(map[string]*pg.Tx),
	}
}

func (ts *Store) GetHandler(ctx context.Context, handler *pg.DB) (orm.DB, error) {
	if !transact.IsTransacted(ctx) {
		return handler, ErrTransactionSQLHandlerNotFound
	}
	return ts.addHandler(ctx, handler)
}

func (ts *Store) addHandler(ctx context.Context, handler *pg.DB) (*pg.Tx, error) {
	ts.mx.Lock()

	storeHandler, ok := ts.handlers[transact.TransactionID(ctx)]
	if !ok {
		storeHandler, err := handler.Begin()
		if err != nil {
			return nil, err
		}
		ts.handlers[transact.TransactionID(ctx)] = storeHandler
	}

	ts.mx.Unlock()

	return storeHandler, nil
}

func (ts *Store) Commit(ctx context.Context) error {
	if transact.IsTransactedChild(ctx) {
		return nil
	}

	ts.mx.Lock()

	storeHandler, ok := ts.handlers[transact.TransactionID(ctx)]
	if !ok {
		ts.mx.Unlock()
		return nil
	}
	var err error
	err = storeHandler.Commit()

	delete(ts.handlers, transact.TransactionID(ctx))
	ts.mx.Unlock()

	return err
}

func (ts *Store) RollBack(ctx context.Context) error {
	if transact.IsTransactedChild(ctx) {
		return nil
	}
	ts.mx.Lock()

	storeHandler, ok := ts.handlers[transact.TransactionID(ctx)]
	if !ok {
		ts.mx.Unlock()
		return nil
	}
	var err error
	err = storeHandler.Rollback()
	delete(ts.handlers, transact.TransactionID(ctx))
	ts.mx.Unlock()

	return err
}
