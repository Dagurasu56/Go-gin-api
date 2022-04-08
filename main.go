package main

import (
	"github.com/Dagurasu56/go-gin-api/database"
	"github.com/Dagurasu56/go-gin-api/routes"
)

func main() {

	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
