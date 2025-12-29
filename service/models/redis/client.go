package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"goliath/utils"
)

var redisUrl string

func init() {
	redisUrl = utils.GetEnv("REDIS_URL", "localhost:6379")
}

func Get(ctx context.Context, key string) (any, error) {
	rdb := getClient()

	var value string
	err := rdb.Get(ctx, key).Scan(&value)

	return value, err
}

func Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	rdb := getClient()
	return rdb.Set(ctx, key, value, expiration).Err()
}

func Del(ctx context.Context, key string) {
	rdb := getClient()
	rdb.Del(ctx, key)
}

func getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "",
		DB:       0,
	})
}
