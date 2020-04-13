package transactor

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/qilin/crm-api/pkg/context/transact"
	"go.uber.org/zap"
)

type Transactor struct {
	TransactorStore *Store
}

func New() *Transactor {
	return &Transactor{
		TransactorStore: newStore(),
	}
}

func (h *Transactor) GetStore() *Store {
	return h.TransactorStore
}

func (h *Transactor) Transact(ctx context.Context, txFunc func(tx context.Context) error) (err error) {
	var tx context.Context

	defer func() {
		if p := recover(); p != nil {
			err = errors.New(fmt.Sprint(p))
		}
		if err != nil {
			errRB := h.TransactorStore.RollBack(tx)
			if errRB != nil {
				zap.Error(errRB)
				return
			}
			return
		}
		err = h.TransactorStore.Commit(tx)
		if err != nil {
			zap.Error(err)
			return
		}

	}()

	if transact.IsTransactedChild(ctx) {
		tx = ctx
	} else {
		tx = transact.BeginTransact(ctx)
	}
	return txFunc(tx)
}
