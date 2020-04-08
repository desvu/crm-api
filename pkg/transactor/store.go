package transactor

import (
	"context"
	"sync"

	"github.com/qilin/crm-api/pkg/context/transact"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type Store struct {
	handlers map[string]sql.TransactionHandler
	mx       sync.RWMutex
}

type Handler interface {
	sql.Handler
	sql.TransactionBeginner
}

func newStore() *Store {
	return &Store{
		handlers: make(map[string]sql.TransactionHandler),
	}
}

func (ts *Store) GetHandler(ctx context.Context, handler Handler) (sql.Handler, bool, error) {
	if !transact.IsTransacted(ctx) {
		return handler, false, nil
	}

	th, err := ts.getTransactHandler(ctx, handler)
	if err != nil {
		return nil, false, err
	}

	return th, true, nil
}

func (ts *Store) getTransactHandler(ctx context.Context, handler sql.TransactionBeginner) (sql.TransactionHandler, error) {
	var err error

	ts.mx.Lock()

	storeHandler, ok := ts.handlers[transact.TransactionID(ctx)]
	if !ok {
		storeHandler, err = handler.Begin()
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
