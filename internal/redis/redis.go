package redis

import (
	"github.com/redis/go-redis/v9"
)

func ConnectCache(config *map[string]string) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     (*config)["REDIS_URL"],
		Username: (*config)["REDIS_USERNAME"],
		Password: (*config)["REDIS_PASSWORD"],
		DB:       0,
	})

	return rdb
}
