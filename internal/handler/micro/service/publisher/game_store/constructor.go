package game_store

import (
	"github.com/isayme/go-amqp-reconnect/rabbitmq"
	"github.com/qilin/crm-api/internal/domain/publisher"
	"github.com/qilin/crm-api/internal/env"
)

const ExchangeName = "crm"
const QueueName = "game_store"

type Publisher struct {
	Channel *rabbitmq.Channel
}

func New(e *env.Env) (publisher.GameStorePublisher, error) {
	ch, err := e.Rabbit.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(QueueName, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		Channel: ch,
	}, nil
}
