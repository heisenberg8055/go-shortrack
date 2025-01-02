package main

import (
	"github.com/heisenberg8055/gotiny/config"
	"github.com/heisenberg8055/gotiny/postgres"
)

func main() {
	env := config.LoadConfig()
	postgres.ConnectDB(&env)
}
