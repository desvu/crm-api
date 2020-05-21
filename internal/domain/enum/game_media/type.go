package game_media

type Type struct {
	ID             uint
	Name           string
	IsNeedResize   bool
	ResultWidth    int
	ResultHeight   int
	IsNeedValidate bool
	AspectHeight   int
	AspectWidth    int
}

func NewType(t uint) Type {
	switch t {
	case 1:
		return Type{
			ID:             1,
			Name:           "wideSlider",
			IsNeedResize:   true,
			ResultWidth:    1064,
			ResultHeight:   599,
			IsNeedValidate: true,
			AspectWidth:    16,
			AspectHeight:   9,
		}
	case 2:
		return Type{
			ID:             2,
			Name:           "vertical",
			IsNeedResize:   true,
			ResultWidth:    200,
			ResultHeight:   266,
			IsNeedValidate: true,
			AspectWidth:    3,
			AspectHeight:   4,
		}
	case 3:
		return Type{
			ID:             3,
			Name:           "horizontal",
			IsNeedResize:   true,
			ResultWidth:    524,
			ResultHeight:   294,
			IsNeedValidate: true,
			AspectWidth:    16,
			AspectHeight:   9,
		}
	case 4:
		return Type{
			ID:             4,
			Name:           "horizontalSmall",
			IsNeedResize:   true,
			ResultWidth:    254,
			ResultHeight:   143,
			IsNeedValidate: true,
			AspectWidth:    16,
			AspectHeight:   9,
		}
	case 5:
		return Type{
			ID:             5,
			Name:           "largeSingle",
			IsNeedResize:   true,
			ResultWidth:    744,
			ResultHeight:   410,
			IsNeedValidate: true,
			AspectWidth:    16,
			AspectHeight:   9,
		}
	case 6:
		return Type{
			ID:             6,
			Name:           "catalog",
			IsNeedResize:   true,
			ResultWidth:    88,
			ResultHeight:   50,
			IsNeedValidate: true,
			AspectWidth:    16,
			AspectHeight:   9,
		}
	case 7:
		return Type{
			ID:             7,
			Name:           "screenshot",
			IsNeedResize:   true,
			ResultWidth:    1064,
			ResultHeight:   562,
			IsNeedValidate: true,
			AspectWidth:    16,
			AspectHeight:   9,
		}
	case 8:
		return Type{
			ID:   8,
			Name: "description",
		}
	default:
		return Type{
			Name: "undefined",
		}
	}
}

func NewTypeByString(t string) Type {
	switch t {
	case "wideSlider":
		return Type{
			ID:             1,
			Name:           "wideSlider",
			IsNeedResize:   true,
			ResultWidth:    1064,
			ResultHeight:   599,
			IsNeedValidate: true,
			AspectWidth:    16,
			AspectHeight:   9,
		}
	case "vertical":
		return Type{
			ID:             2,
			Name:           "vertical",
			IsNeedResize:   true,
			ResultWidth:    200,
			ResultHeight:   266,
			IsNeedValidate: true,
			AspectWidth:    3,
			AspectHeight:   4,
		}
	case "horizontal":
		return Type{
			ID:             3,
			Name:           "horizontal",
			IsNeedResize:   true,
			ResultWidth:    524,
			ResultHeight:   294,
			IsNeedValidate: true,
			AspectWidth:    16,
			AspectHeight:   9,
		}
	case "horizontalSmall":
		return Type{
			ID:             4,
			Name:           "horizontalSmall",
			IsNeedResize:   true,
			ResultWidth:    254,
			ResultHeight:   143,
			IsNeedValidate: true,
			AspectWidth:    16,
			AspectHeight:   9,
		}
	case "largeSingle":
		return Type{
			ID:             5,
			Name:           "largeSingle",
			IsNeedResize:   true,
			ResultWidth:    744,
			ResultHeight:   410,
			IsNeedValidate: true,
			AspectWidth:    16,
			AspectHeight:   9,
		}
	case "catalog":
		return Type{
			ID:             6,
			Name:           "catalog",
			IsNeedResize:   true,
			ResultWidth:    88,
			ResultHeight:   50,
			IsNeedValidate: true,
			AspectWidth:    16,
			AspectHeight:   9,
		}
	case "screenshot":
		return Type{
			ID:             7,
			Name:           "screenshot",
			IsNeedResize:   true,
			ResultWidth:    1064,
			ResultHeight:   562,
			IsNeedValidate: true,
			AspectWidth:    16,
			AspectHeight:   9,
		}
	case "description":
		return Type{
			ID:   8,
			Name: "description",
		}
	default:
		return Type{
			Name: "undefined",
		}
	}
}
