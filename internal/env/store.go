package env

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/go-redis/redis/v7"
	"github.com/qilin/crm-api/internal/config"
)

type Store struct {
	Postgres *Postgres
	Redis    *Redis
}

type Postgres struct {
	Conn *pg.DB
}

type Redis struct {
	Client *redis.Client
}

func newStore(ctx context.Context, conf config.Store) (*Store, error) {
	postgresEnv, err := newPostgres(ctx, conf.Mongo)
	if err != nil {
		return nil, err
	}

	redisEnv, err := newRedis(ctx, conf.Redis)
	if err != nil {
		return nil, err
	}

	return &Store{
		Postgres: postgresEnv,
		Redis:    redisEnv,
	}, nil
}

func newPostgres(ctx context.Context, conf config.PostgresConf) (*Postgres, error) {
	postgres := &Postgres{}
	postgres.Conn = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Database: conf.Database,
		User:     conf.User,
		Password: conf.Password,
	})

	return postgres, nil
}

func newRedis(ctx context.Context, conf config.RedisConf) (*Redis, error) {
	//client := redis.NewClient(&redis.Options{
	//	Addr:     conf.Host,
	//	Password: conf.Password,
	//	DB:       conf.DB,
	//})
	//
	//_, err := client.Ping().Result()
	//if err != nil {
	//	return nil, err
	//}

	return &Redis{nil}, nil
}
