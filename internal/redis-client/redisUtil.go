package redis_client

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	log_middleware "github.com/heisenberg8055/gotiny/internal/log"
	"github.com/redis/go-redis/v9"
)

func RedisSet(r *http.Request, client *redis.Client, longURL string, shortURL string, logger *slog.Logger) bool {
	currTime := time.Now()
	err := client.Set(context.Background(), shortURL, longURL, 0).Err()
	elapsed := time.Since(currTime)
	if err != nil {
		log_middleware.LogError(log_middleware.Response{Method: r.Method, Url: r.URL.Path, Status: http.StatusInternalServerError, Message: err.Error(), TimeTaken: elapsed.String()}, logger, "Failed to Update Cache for shortURL: "+shortURL)
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
