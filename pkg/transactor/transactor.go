package transactor

import (
	"context"
	"errors"
	"fmt"
	"log"
	"runtime"
	"runtime/debug"

	"github.com/qilin/crm-api/pkg/context/transact"
)

func New() (*Transactor, *Store) {
	store := newStore()
	return &Transactor{
		TransactorStore: store,
	}, store
}

type Transactor struct {
	TransactorStore *Store
}

func (h *Transactor) Transact(ctx context.Context, txFunc func(tx context.Context) error) (err error) {
	var tx context.Context

	defer func() {
		if p := recover(); p != nil {
			debug.PrintStack()
			_, file, line, _ := runtime.Caller(1)
			log.Println(file, line, p)
			err = errors.New(fmt.Sprint(p))
		}
		if err != nil {
			errRB := h.TransactorStore.RollBack(tx)
			if errRB != nil {
				return
			}
			return
		}
		err = h.TransactorStore.Commit(tx)
		if err != nil {
			log.Println(err)
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
