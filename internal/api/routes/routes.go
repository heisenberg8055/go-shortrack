package routes

import (
	"log/slog"
	"net/http"

	"github.com/heisenberg8055/gotiny/internal/api/middleware"
	handlers "github.com/heisenberg8055/gotiny/internal/api/routes/handler"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func Routes(postClient *pgxpool.Pool, redisClient *redis.Client, logger *slog.Logger) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", middleware.Middleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.Home(w, r, logger)
	}, logger))

	router.HandleFunc("POST /", middleware.Middleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.AddURL(w, r, postClient, redisClient, logger)
	}, logger))

	router.HandleFunc("GET /{shortUrl}", middleware.Middleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetURL(w, r, postClient, redisClient, logger)
	}, logger))

	router.HandleFunc("GET /count", middleware.Middleware(func(w http.ResponseWriter, r *http.Request) {
		handlers.GetCount(w, r, postClient, logger)
	}, logger))

	router.Handle("GET /healthz", &handlers.Health{})

	return router
}
