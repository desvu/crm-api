package sub_game_store

import (
	"github.com/micro/go-micro/v2/broker"
)

const Topic = "game_store"

type Subscriber struct {
	sub broker.Subscriber
}

func New(b broker.Broker) (*Subscriber, error) {
	var err error
	subscriber := new(Subscriber)

	subscriber.sub, err = b.Subscribe(Topic, subscriber.Handler)
	if err != nil {
		return nil, err
	}

	return subscriber, nil
}
