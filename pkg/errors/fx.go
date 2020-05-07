package errors

import (
	"errors"
	"strings"
)

func WrapFxError(err error) error {
	msg := strings.Split(err.Error(), "): ")
	return NewInternal(errors.New(msg[len(msg)-1]))
}
