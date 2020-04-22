package game_store

import (
	"github.com/micro/go-micro/v2/broker"
	"github.com/qilin/crm-api/internal/domain/publisher"
)

type Data struct {
	GameStorePublishID uint
	GameID             string
	Body               string
}

func (p Publisher) Publish(data publisher.PublishGameStoreData) error {
	str := "{\"game_id\":\"123\"}"
	return p.broker.Publish(Topic, &broker.Message{
		Body: []byte(str),
	})
}
