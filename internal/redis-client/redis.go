package redis_client

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func ConnectCache() *redis.Client {
	rdb, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatal(err.Error())
		return nil
	}
	return redis.NewClient(rdb)
}
