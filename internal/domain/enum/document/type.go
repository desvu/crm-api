package document

type Type uint8

const (
	TypeUndefined Type = iota
	TypeLicense
	TypeGameRules
	TypePrivacyPolicy
)

const (
	undefined     = "undefined"
	license       = "license"
	gameRules     = "game_rules"
	privacyPolicy = "privacy_policy"
)

func NewType(v uint8) Type {
	var source Type

	switch v := Type(v); v {
	case TypeLicense,
		TypeGameRules,
		TypePrivacyPolicy:
		source = v
	}

	return source
}

func NewTypeByString(v string) Type {
	switch v {
	case license:
		return TypeLicense
	case gameRules:
		return TypeGameRules
	case privacyPolicy:
		return TypePrivacyPolicy
	default:
		return TypeUndefined
	}
}

func NewTypePointerByStringPointer(v *string) *Type {
	if v == nil {
		return nil
	}

	t := NewTypeByString(*v)
	return &t
}

func (t Type) Value() uint8 {
	return uint8(t)
}

func (t Type) String() string {
	switch t {
	case TypeLicense:
		return license
	case TypeGameRules:
		return gameRules
	case TypePrivacyPolicy:
		return privacyPolicy
	default:
		return undefined
	}
}
