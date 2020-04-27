package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/qilin/crm-api/internal/auth"
)

type Config struct {
	Store Store
	Auth  auth.Config
}

type Store struct {
	Postgres PostgresConf
	Redis    RedisConf
	Rabbit   RabbitConf
}

type PostgresConf struct {
	Host     string `default:"127.0.0.1"`
	Database string `default:"crm"`
	Port     string `default:"5432"`
	User     string `default:"postgres"`
	Password string `default:"password"`
}

type RedisConf struct {
	Host     string `default:"127.0.0.1:6379"`
	Password string `default:""`
	DB       int    `default:"0"`
}

type RabbitConf struct {
	Url string `default:"amqp://guest:guest@localhost:5672"`
}

func New() (*Config, error) {
	var c Config
	if err := envconfig.Process("qilin", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
