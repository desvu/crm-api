package entity

type Review struct {
	ID             uint
	GameRevisionID uint
	PressName      string
	Link           string
	Score          uint
	Quote          string
}
