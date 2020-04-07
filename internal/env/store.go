package env

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/go-redis/redis/v7"
	"github.com/qilin/crm-api/internal/config"
	"github.com/qilin/crm-api/internal/env/transaction"
	"github.com/qilin/crm-api/pkg/context/transact"
)

type Store struct {
	Postgres *Postgres
	Redis    *Redis
}

type Postgres struct {
	Handler          *pg.DB
	TransactionStore *transaction.Store
}

type Redis struct {
	Client *redis.Client
}

func newStore(ctx context.Context, conf config.Store) (*Store, error) {
	postgresEnv, err := newPostgres(ctx, conf.Postgres)
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
	postgres := &Postgres{
		TransactionStore: transaction.New(),
	}
	postgres.Handler = pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Database: conf.Database,
		User:     conf.User,
		Password: conf.Password,
	})

	return postgres, nil
}

func (p Postgres) GetHandler(ctx context.Context) (orm.DB, error) {
	if !transact.IsTransacted(ctx) {
		return p.Handler, nil
	}

	return p.TransactionStore.GetHandler(ctx, p.Handler)
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
