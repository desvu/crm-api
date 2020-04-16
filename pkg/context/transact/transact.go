package transact

import (
	"context"

	"github.com/google/uuid"
)

type key int

const transactKey key = 0

func IsTransacted(ctx context.Context) bool {
	_, ok := FromContext(ctx)
	return ok
}

func TransactionID(ctx context.Context) string {
	t, ok := FromContext(ctx)
	if !ok {
		return ""
	}
	return t.TransactionID()
}

func BeginTransact(ctx context.Context) context.Context {
	t, ok := FromContext(ctx)
	if !ok {
		return newContext(ctx, newTransact())
	}
	return newContext(ctx, newChildTransact(t.id))
}

func IsTransactedChild(ctx context.Context) bool {
	t, ok := FromContext(ctx)
	if !ok {
		return false
	}
	return t.isChild
}

func newContext(ctx context.Context, transact *transact) context.Context {
	return context.WithValue(ctx, transactKey, transact)
}

func FromContext(ctx context.Context) (*transact, bool) {
	device, ok := ctx.Value(transactKey).(*transact)
	return device, ok
}

func newTransact() *transact {
	return &transact{
		id: uuid.New().String(),
	}
}

func newChildTransact(id string) *transact {
	return &transact{
		id:      id,
		isChild: true,
	}
}

type transact struct {
	id      string
	isChild bool
}

func (t *transact) TransactionID() string {
	return t.id
}

func (t *transact) IsTransactedChild() bool {
	return t.isChild
}
