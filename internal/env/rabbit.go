package env

import (
	"github.com/isayme/go-amqp-reconnect/rabbitmq"
)

func newRabbit() (*rabbitmq.Connection, error) {
	return rabbitmq.Dial("amqp://guest:guest@localhost:5672")
}
