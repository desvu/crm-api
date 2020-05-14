package block

//swagger:enum Type
type Type string

const (
	TypeWideSlider    Type = "wide_slider"
	TypeBasicLayout1  Type = "basic_layout_1"
	TypeBasicLayout2  Type = "basic_layout_2"
	TypeBasicLayout4  Type = "basic_layout_4"
	TypeBasicLayout5  Type = "basic_layout_5"
	TypeBasicLayout42 Type = "basic_layout_4_2"
	TypeCatalogue10   Type = "catalog_10"
	TypeCatalogue20   Type = "catalog_20"
)

func (t Type) String() string {
	return string(t)
}

func (t Type) Valid() bool {
	switch t {
	case
		TypeWideSlider,
		TypeBasicLayout1,
		TypeBasicLayout2,
		TypeBasicLayout4,
		TypeBasicLayout5,
		TypeBasicLayout42,
		TypeCatalogue10,
		TypeCatalogue20:
		return true
	}
	return false
}

func (t Type) ValidGamesCount(n int) bool {
	switch t {
	case TypeWideSlider:
		return 3 <= n && n <= 8
	case TypeBasicLayout1:
		return n == 1
	case TypeBasicLayout2:
		return n%2 == 0
	case TypeBasicLayout4:
		return n%4 == 0
	case TypeBasicLayout5:
		return n%5 == 0
	case TypeBasicLayout42:
		return n%8 == 0
	case TypeCatalogue10:
		return n == 10
	case TypeCatalogue20:
		return n == 20
	}
	return false
}
