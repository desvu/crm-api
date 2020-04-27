package env

import (
	"github.com/isayme/go-amqp-reconnect/rabbitmq"
	"github.com/qilin/crm-api/internal/config"
)

func newRabbit(cfg config.RabbitConf) (*rabbitmq.Connection, error) {
	return rabbitmq.Dial(cfg.Uri)
}
