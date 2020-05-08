package block

type Type string

const (
	Type_WideSlider    Type = "wide_slider"
	Type_BasicLayout1  Type = "basic_layout_1"
	Type_BasicLayout2  Type = "basic_layout_2"
	Type_BasicLayout4  Type = "basic_layout_4"
	Type_BasicLayout5  Type = "basic_layout_5"
	Type_BasicLayout42 Type = "basic_layout_4_2"
	Type_Catalogue10   Type = "catalog_10"
	Type_Catalogue20   Type = "catalog_20"
)

func (t Type) Valid() bool {
	switch t {
	case
		Type_WideSlider,
		Type_BasicLayout1,
		Type_BasicLayout2,
		Type_BasicLayout4,
		Type_BasicLayout5,
		Type_BasicLayout42,
		Type_Catalogue10,
		Type_Catalogue20:
		return true
	}
	return false
}

func (t Type) ValidGamesCount(n int) bool {
	switch t {
	case Type_WideSlider:
		return 3 <= n && n <= 8
	case Type_BasicLayout1:
		return n == 1
	case Type_BasicLayout2:
		return n%2 == 0
	case Type_BasicLayout4:
		return n%4 == 0
	case Type_BasicLayout5:
		return n%5 == 0
	case Type_BasicLayout42:
		return n%8 == 0
	case Type_Catalogue10:
		return n == 10
	case Type_Catalogue20:
		return n == 20
	}
	return false
}
