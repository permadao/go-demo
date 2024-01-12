package redis

import (
	"context"
	"time"

	"github.com/permadao/go-demo/storage/redis/schema"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	rdb *redis.Client
	ctx context.Context
}

func New(redisURL string) *Redis {
	redisOpt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}
	redisOpt.DialTimeout = 200 * time.Second
	redisOpt.PoolSize = 1000

	return &Redis{
		rdb: redis.NewClient(redisOpt),
		ctx: context.Background(),
	}
}

func (r *Redis) SetTest(key, value string) error {
	return r.rdb.Set(r.ctx, schema.RdTest+key, value, 0).Err()
}

func (r *Redis) GetTest(key string) (string, error) {
	return r.rdb.Get(r.ctx, schema.RdTest+key).Result()
}
