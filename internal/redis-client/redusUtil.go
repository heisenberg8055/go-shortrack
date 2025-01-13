package redis_client

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func RedisSet(client *redis.Client, longURL string, shortURL string) bool {
	err := client.Set(context.Background(), longURL, shortURL, 0).Err()
	if err != nil {
		log.Printf("Failed to update cache: %v", err)
		return false
	}
	return true
}

func RedisGet(client *redis.Client, longURL string) string {
	shortURL, err := client.Get(context.Background(), longURL).Result()
	if err == redis.Nil {
		return ""
	}
	return shortURL
}
