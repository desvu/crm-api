package grpcerror

import (
	e "github.com/pkg/errors"
	"github.com/qilin/crm-api/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(err error) error {
	var resultErr error
	switch v := e.Cause(err).(type) {
	case errors.Error:
		resultErr = status.Error(getCodeByType(v.Type), err.Error())
	default:
		resultErr = status.Error(codes.Internal, "oops, something went wrong...")
	}

	zap.L().Error(err.Error())

	return resultErr
}

func getCodeByType(t errors.Type) codes.Code {
	switch t {
	case errors.ErrNotFound:
		return codes.NotFound
	case errors.ErrAlreadyExist:
		return codes.AlreadyExists
	case errors.ErrValidation:
		return codes.InvalidArgument
	case errors.ErrInternal:
		return codes.Internal
	default:
		return codes.Internal
	}
}
