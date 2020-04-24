package publisher

import (
	"time"

	"github.com/qilin/crm-api/internal/domain/entity"
)

type GameStorePublisher interface {
	Publish(data PublishGameStoreData) error
}

type PublishGameStoreData struct {
	GameStorePublish *entity.GameStorePublish
	Game             *entity.GameEx
}

type GameStoreMessage struct {
	ID          string
	Title       string
	Type        string
	RevisionID  uint
	Summary     string
	Description string
	Slug        string
	License     string
	Platforms   []string
	ReleaseDate time.Time

	Tags       []GameStoreTagMessage
	Features   []GameStoreFeatureMessage
	Developers []GameStoreDeveloperMessage
	Publishers []GameStorePublisherMessage
	Genres     []GameStoreGenreMessage
	Rand       string
}

type GameStoreTagMessage struct {
	ID   uint
	Name string
}

type GameStoreFeatureMessage struct {
	ID   uint
	Name string
}

type GameStoreDeveloperMessage struct {
	ID   uint
	Name string
}

type GameStorePublisherMessage struct {
	ID   uint
	Name string
}

type GameStoreGenreMessage struct {
	ID   uint
	Name string
}
