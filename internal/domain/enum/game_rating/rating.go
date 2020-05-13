package game_rating

import "strings"

type Rating uint8

const (
	RatingUndefined = iota
	BBFC_U
	BBFC_PG
	BBFC_12A
	BBFC_12
	BBFC_15
	BBFC_18
	BBFC_R18

	CERO_A
	CERO_B
	CERO_C
	CERO_D
	CERO_Z

	ESRB_EC
	ESRB_E
	ESRB_E10Plus
	ESRB_T
	ESRB_M
	ESRB_A
	ESRB_RP

	PEGI_3
	PEGI_7
	PEGI_12
	PEGI_16
	PEGI_18

	USK_USK
	USK_0
	USK_6
	USK_12
	USK_16
	USK_18
)

func NewRating(v uint8) Rating {
	var source Rating

	if v >= BBFC_U && v <= USK_18 {
		source = Rating(v)
	}

	return source
}

func NewRatingByString(agency, rating string) Rating {
	switch strings.ToUpper(agency) {
	case "BBFC":
		switch strings.ToUpper(rating) {
		case "U":
			return BBFC_U
		case "PG":
			return BBFC_PG
		case "12A":
			return BBFC_12A
		case "12":
			return BBFC_12
		case "15":
			return BBFC_15
		case "18":
			return BBFC_18
		case "R18":
			return BBFC_R18
		}
	case "CERO":
		switch strings.ToUpper(rating) {
		case "A":
			return CERO_A
		case "B":
			return CERO_B
		case "C":
			return CERO_C
		case "D":
			return CERO_D
		case "Z":
			return CERO_Z
		}
	case "ESRB":
		switch strings.ToUpper(rating) {
		case "EC":
			return ESRB_EC
		case "E":
			return ESRB_E
		case "E10+":
			return ESRB_E10Plus
		case "T":
			return ESRB_T
		case "M":
			return ESRB_T
		case "A":
			return ESRB_A
		case "RP":
			return ESRB_RP
		}
	case "PEGI":
		switch strings.ToUpper(rating) {
		case "3":
			return PEGI_3
		case "7":
			return PEGI_7
		case "12":
			return PEGI_12
		case "16":
			return PEGI_16
		case "18":
			return PEGI_18
		}
	case "USK":
		switch strings.ToUpper(rating) {
		case "USK":
			return USK_USK
		case "0":
			return USK_0
		case "6":
			return USK_6
		case "12":
			return USK_12
		case "16":
			return USK_16
		case "18":
			return USK_18
		}
	}
	return RatingUndefined
}

func (r Rating) Value() uint8 {
	return uint8(r)
}

func (r Rating) String() string {
	switch r {
	case BBFC_U:
		return "U"
	case BBFC_PG:
		return "PG"
	case BBFC_12A:
		return "12A"
	case BBFC_12:
		return "12"
	case BBFC_15:
		return "15"
	case BBFC_18:
		return "18"
	case BBFC_R18:
		return "R18"

	case CERO_A:
		return "A"
	case CERO_B:
		return "B"
	case CERO_C:
		return "C"
	case CERO_D:
		return "D"
	case CERO_Z:
		return "Z"

	case ESRB_EC:
		return "EC"
	case ESRB_E:
		return "E"
	case ESRB_E10Plus:
		return "E10+"
	case ESRB_T:
		return "T"
	case ESRB_M:
		return "M"
	case ESRB_A:
		return "A"
	case ESRB_RP:
		return "RP"

	case PEGI_3:
		return "3"
	case PEGI_7:
		return "7"
	case PEGI_12:
		return "12"
	case PEGI_16:
		return "16"
	case PEGI_18:
		return "18"

	case USK_USK:
		return "USK"
	case USK_0:
		return "0"
	case USK_6:
		return "6"
	case USK_12:
		return "12"
	case USK_16:
		return "16"
	case USK_18:
		return "18"
	default:
		return "undefined"
	}
}
