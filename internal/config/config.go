package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App     AppConf
	Store   Store
	Rabbit  RabbitConf
	Storage StorageConf
}

type AppConf struct {
	StorageURL string `envconfig:"storage_url"`
}

type Store struct {
	Postgres PostgresConf
	Redis    RedisConf
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
	Uri string `default:"amqp://guest:guest@localhost:5672"`
}

type StorageConf struct {
	Bucket string
}

func New() (*Config, error) {
	var c Config
	if err := envconfig.Process("qilin", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
