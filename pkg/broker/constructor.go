package broker

import (
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/broker/nats"
)

func New() (broker.Broker, error) {
	b := nats.NewBroker()

	if err := b.Init(); err != nil {
		return nil, err
	}

	if err := b.Connect(); err != nil {
		return nil, err
	}

	return b, nil
}
