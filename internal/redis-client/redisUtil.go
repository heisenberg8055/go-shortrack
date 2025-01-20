package redis_client

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func RedisSet(client *redis.Client, longURL string, shortURL string) bool {
	err := client.Set(context.Background(), shortURL, longURL, 0).Err()
	if err != nil {
		log.Printf("Failed to update cache: %v", err)
		return false
	}
	return true
}

func RedisGet(client *redis.Client, shortURL string) string {
	shortURL, err := client.Get(context.Background(), shortURL).Result()
	if err == redis.Nil {
		return ""
	}
	return shortURL
}
