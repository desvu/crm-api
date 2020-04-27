package errors

type Type uint8

const (
	ErrNotFound     Type = 1
	ErrAlreadyExist Type = 2
	ErrValidation   Type = 3
	ErrInternal     Type = 4
)

func (t Type) String() string {
	switch t {
	case ErrNotFound:
		return "not found"
	case ErrAlreadyExist:
		return "already exist"
	case ErrValidation:
		return "validation"
	case ErrInternal:
		return "internal"
	default:
		return "internal"
	}
}
