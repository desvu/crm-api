package errors

import (
	"github.com/pkg/errors"
)

type Error struct {
	Err   error
	Key   string
	Type  Type
	Layer Layer
}

func NewService(t Type, msg string, key string) Error {
	return Error{
		Err:   errors.WithStack(errors.New(msg)),
		Key:   key,
		Type:  t,
		Layer: Service,
	}
}

func NewAPI(t Type, msg string, key string) Error {
	return Error{
		Err:   errors.WithStack(errors.New(msg)),
		Key:   key,
		Type:  t,
		Layer: API,
	}
}

func NewInternal(err error) Error {
	return Error{
		Err:   errors.WithStack(err),
		Key:   "internal_server_error",
		Type:  ErrInternal,
		Layer: Internal,
	}
}

func (e Error) Error() string {
	return e.Err.Error()
}
