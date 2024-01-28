package main

import (
	"github.com/rodriguesabner/ifinance-back/config"
	"github.com/rodriguesabner/ifinance-back/database"
	"github.com/rodriguesabner/ifinance-back/router"
)

func main() {
	config.LoadEnv()
	port := config.GetPort()

	database.ConnectDB()
	r := router.SetupRouter()

	StartServer(port, r)
}
