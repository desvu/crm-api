package game_media

type Type uint

const (
	TypeUndefined       Type = 0
	TypeWideSlider      Type = 1
	TypeVertical        Type = 2
	TypeHorizontal      Type = 3
	TypeHorizontalSmall Type = 4
	TypeLargeSingle     Type = 5
	TypeCatalog         Type = 6
	TypeScreenshot      Type = 7
	TypeDescription     Type = 8
)

func NewType(t uint) Type {
	switch t {
	case 1:
		return TypeWideSlider
	case 2:
		return TypeVertical
	case 3:
		return TypeHorizontal
	case 4:
		return TypeHorizontalSmall
	case 5:
		return TypeLargeSingle
	case 6:
		return TypeCatalog
	case 7:
		return TypeScreenshot
	case 8:
		return TypeDescription
	default:
		return TypeUndefined
	}
}

func NewTypeByString(t string) Type {
	switch t {
	case "wideSlider":
		return TypeWideSlider
	case "vertical":
		return TypeVertical
	case "horizontal":
		return TypeHorizontal
	case "horizontalSmall":
		return TypeHorizontalSmall
	case "largeSingle":
		return TypeLargeSingle
	case "catalog":
		return TypeCatalog
	case "screenshot":
		return TypeScreenshot
	case "description":
		return TypeDescription
	default:
		return TypeUndefined
	}
}

func (t Type) Value() uint {
	return uint(t)
}

func (t Type) String() string {
	switch t {
	case TypeWideSlider:
		return "wideSlider"
	case TypeVertical:
		return "vertical"
	case TypeHorizontal:
		return "horizontal"
	case TypeHorizontalSmall:
		return "horizontalSmall"
	case TypeLargeSingle:
		return "largeSingle"
	case TypeCatalog:
		return "catalog"
	case TypeScreenshot:
		return "screenshot"
	case TypeDescription:
		return "description"
	default:
		return "undefined"
	}
}
