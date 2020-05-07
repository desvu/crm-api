package game

type Type uint8

const (
	TypeUndefined Type = 0
	TypeWeb       Type = 1
	TypeDesktop   Type = 2
)

func NewType(v uint8) Type {
	var source Type

	switch v := Type(v); v {
	case TypeWeb,
		TypeDesktop:
		source = v
	}

	return source
}

func NewTypeByString(v string) Type {
	switch v {
	case "web":
		return TypeWeb
	case "desktop":
		return TypeDesktop
	default:
		return TypeUndefined
	}
}

func NewTypePointerByStringPointer(v *string) *Type {
	if v == nil {
		return nil
	}

	var t = TypeUndefined
	switch *v {
	case "web":
		t = TypeWeb
		return &t
	case "desktop":
		t = TypeDesktop
		return &t
	default:
		return &t
	}
}

func (t Type) Value() uint8 {
	return uint8(t)
}

func (t Type) String() string {
	switch t {
	case TypeWeb:
		return "web"
	case TypeDesktop:
		return "desktop"
	default:
		return "undefined"
	}
}
