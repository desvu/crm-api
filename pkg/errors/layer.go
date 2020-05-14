package errors

type Layer uint8

const (
	Service  Layer = 1
	Internal Layer = 2
	API      Layer = 3
)

func (l Layer) String() string {
	switch l {
	case Service:
		return "service"
	case API:
		return "api"
	case Internal:
		return "internal"
	default:
		return "internal"
	}
}
