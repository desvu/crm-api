package errors

import (
	"github.com/pkg/errors"
)

type Error struct {
	Err   error
	Type  Type
	Layer Layer
}

func NewService(t Type, msg string) Error {
	return Error{
		Err:   errors.WithStack(errors.New(msg)),
		Type:  t,
		Layer: Service,
	}
}

func NewInternal(err error) Error {
	return Error{
		Err:   errors.WithStack(err),
		Type:  ErrInternal,
		Layer: Internal,
	}
}

func (e Error) Error() string {
	return e.Err.Error()
}
