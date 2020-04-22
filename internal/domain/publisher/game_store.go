package publisher

type GameStorePublisher interface {
	Publish(data PublishGameStoreData) error
}

type PublishGameStoreData struct {
	GameStorePublishID uint
	GameID             string
	Body               string
}
