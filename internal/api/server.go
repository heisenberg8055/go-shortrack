package api

import (
	"log"
	"net/http"

	"github.com/heisenberg8055/gotiny/config"
	"github.com/heisenberg8055/gotiny/internal/api/routes"
	"github.com/heisenberg8055/gotiny/internal/postgres"
	"github.com/heisenberg8055/gotiny/internal/redis"
)

func StartServer() {
	env := config.LoadConfig()
	postClient := postgres.ConnectDB(&env)
	redisClient := redis.ConnectCache(&env)
	mux := routes.Routes(postClient, redisClient)
	server := http.Server{
		Addr:     ":8080",
		Handler:  mux,
		ErrorLog: log.Default(),
	}
	server.ListenAndServe()
}
