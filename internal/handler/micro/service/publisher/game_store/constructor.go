package game_store

import (
	"github.com/micro/go-micro/v2/broker"
	"github.com/qilin/crm-api/internal/domain/publisher"
)

const Topic = "game_store"

type Publisher struct {
	broker broker.Broker
}

func New(b broker.Broker) publisher.GameStorePublisher {
	return &Publisher{
		broker: b,
	}
}
