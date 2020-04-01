package game

type Type uint8

const (
	TypeWeb     Type = 1
	TypeDesktop Type = 2
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

func (t Type) Value() uint8 {
	return uint8(t)
}

func (t Type) String() string {
	var str string

	switch t {
	case TypeWeb:
		str = "web"
	case TypeDesktop:
		str = "desktop"
	}

	return str
}
