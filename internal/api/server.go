package api

import (
	"log"
	"net/http"
	"os"

	"github.com/heisenberg8055/gotiny/config"
	"github.com/heisenberg8055/gotiny/internal/api/routes"
	log_middleware "github.com/heisenberg8055/gotiny/internal/log"
	"github.com/heisenberg8055/gotiny/internal/postgres"
	redis_client "github.com/heisenberg8055/gotiny/internal/redis-client"
)

func StartServer() {
	config.LoadConfig()
	postClient, _ := postgres.ConnectDB()
	redisClient := redis_client.ConnectCache()
	logger := log_middleware.NewLogger()
	mux := routes.Routes(postClient.Db, redisClient, logger)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := http.Server{
		Addr:     ":" + port,
		Handler:  mux,
		ErrorLog: log.Default(),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
