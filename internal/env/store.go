package env

import (
	"github.com/qilin/crm-api/pkg/transactor"

	"github.com/go-pg/pg/v9"
	"github.com/go-redis/redis/v7"
	"github.com/qilin/crm-api/internal/config"
	"github.com/qilin/crm-api/pkg/repository/handler/postgres"
	"github.com/qilin/crm-api/pkg/repository/handler/sql"
)

type Store struct {
	Postgres *Postgres
	Redis    *Redis
}

type Postgres struct {
	Handler    sql.Handler
	Connection *pg.DB
}

type Redis struct {
	Client *redis.Client
}

func newStore(conf config.Store, transactionStore *transactor.Store) (*Store, error) {
	postgresEnv, err := newPostgres(conf.Postgres, transactionStore)
	if err != nil {
		return nil, err
	}

	redisEnv, err := newRedis(conf.Redis)
	if err != nil {
		return nil, err
	}

	return &Store{
		Postgres: postgresEnv,
		Redis:    redisEnv,
	}, nil
}

func newPostgres(conf config.PostgresConf, transactionStore *transactor.Store) (*Postgres, error) {
	return &Postgres{
		Handler: postgres.New(
			postgres.Config{
				Host:     conf.Host,
				Port:     conf.Port,
				Database: conf.Database,
				User:     conf.User,
				Password: conf.Password,
			},
			transactionStore,
		),
	}, nil
}

func newRedis(conf config.RedisConf) (*Redis, error) {
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
