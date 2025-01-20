package api

import (
	"log"
	"net/http"

	"github.com/heisenberg8055/gotiny/config"
	"github.com/heisenberg8055/gotiny/internal/api/routes"
	"github.com/heisenberg8055/gotiny/internal/postgres"
	redis_client "github.com/heisenberg8055/gotiny/internal/redis-client"
)

func StartServer() {
	env := config.LoadConfig()
	postClient, _ := postgres.ConnectDB(&env)
	redisClient := redis_client.ConnectCache(&env)
	mux := routes.Routes(postClient.Db, redisClient)
	server := http.Server{
		Addr:     ":8080",
		Handler:  mux,
		ErrorLog: log.Default(),
	}
	server.ListenAndServe()
}
