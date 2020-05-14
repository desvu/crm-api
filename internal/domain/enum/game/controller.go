package game

type Controller uint8

const (
	ControllerUndefined Controller = iota
	ControllerNotSupported
	ControllerPartiallySupported
	ControllerFullSupport
)

func NewController(v uint8) Controller {
	var source Controller

	switch v := Controller(v); v {
	case ControllerNotSupported,
		ControllerPartiallySupported,
		ControllerFullSupport:
		source = v
	}

	return source
}

func NewControllerByString(v string) Controller {
	switch v {
	case "not":
		return ControllerNotSupported
	case "partially":
		return ControllerPartiallySupported
	case "full":
		return ControllerFullSupport
	default:
		return ControllerUndefined
	}
}

func NewControllerPointerByStringPointer(v *string) *Controller {
	if v == nil {
		return nil
	}
	t := NewControllerByString(*v)
	return &t
}

func (t Controller) Value() uint8 {
	return uint8(t)
}

func (t Controller) String() string {
	switch t {
	case ControllerNotSupported:
		return "not"
	case ControllerPartiallySupported:
		return "partially"
	case ControllerFullSupport:
		return "full"
	default:
		return "undefined"
	}
}
