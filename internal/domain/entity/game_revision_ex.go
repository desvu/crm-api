package entity

type GameRevisionEx struct {
	GameRevision
	Tags       []Tag
	Developers []Developer
	Publishers []Publisher
	Features   []Feature
	Genres     []Genre
}
