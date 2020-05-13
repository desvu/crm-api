package game_rating

import "strings"

type Agency uint8

const (
	AgencyUndefined Agency = iota
	PEGI
	ESRB
	BBFC
	USK
	CERO
)

func NewAgency(v uint8) Agency {
	var source Agency

	switch v := Agency(v); v {
	case AgencyUndefined, PEGI, ESRB, BBFC, USK, CERO:
		source = v
	}

	return source
}

func NewAgencyByString(v string) Agency {
	switch strings.ToUpper(v) {
	case "PEGI":
		return PEGI
	case "ESRB":
		return ESRB
	case "BBFC":
		return BBFC
	case "USK":
		return USK
	case "CERO":
		return CERO
	default:
		return AgencyUndefined
	}
}

func (t Agency) Value() uint8 {
	return uint8(t)
}

func (t Agency) String() string {
	switch t {
	case PEGI:
		return "PEGI"
	case ESRB:
		return "ESRB"
	case BBFC:
		return "BBFC"
	case USK:
		return "USK"
	case CERO:
		return "CERO"
	default:
		return "undefined"
	}
}
