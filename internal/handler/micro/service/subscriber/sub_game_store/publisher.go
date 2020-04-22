package sub_game_store

import (
	"log"

	"github.com/micro/go-micro/v2/broker"
)

type Data struct {
	GameStorePublishID uint
	GameID             string
	Body               string
}

func (p Subscriber) Handler(e broker.Event) error {
	if e.Error() != nil {
		return e.Error()
	}

	log.Println(e.Topic())
	log.Println(string(e.Message().Body))

	return nil
}
