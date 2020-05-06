package entity

type Localization struct {
	ID             uint
	GameRevisionID uint
	Language       string
	Interface      bool
	Audio          bool
	Subtitles      bool
}
