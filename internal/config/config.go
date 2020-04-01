package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Store Store
}

type Store struct {
	Postgres PostgresConf `envconfig:"POSTGRES"`
	Redis    RedisConf    `envconfig:"REDIS"`
}

type PostgresConf struct {
	Host     string `envconfig:"P_HOST" required:"false" default:"127.0.0.1"`
	Database string `envconfig:"P_DATABASE" required:"false" default:"crm"`
	Port     string `envconfig:"P_PORT" required:"false" default:"5432"`
	User     string `envconfig:"P_USER" required:"false" default:"postgres"`
	Password string `envconfig:"P_PASSWORD" required:"false" default:"password"`
}

type RedisConf struct {
	Host     string `envconfig:"R_HOST" required:"false" default:"127.0.0.1:6379"`
	Password string `envconfig:"R_PASSWORD" required:"false" default:""`
	DB       int    `envconfig:"R_DB" required:"false" default:"0"`
}

func New() (*Config, error) {
	var c Config
	if err := envconfig.Process("AUTHONE", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
