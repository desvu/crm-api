package entity

type Review struct {
	ID             uint
	GameRevisionID uint
	PressName      string
	Link           string
	Score          string
	Quote          string
}
