package main

import (
	"github.com/ivancl4udio/gin-api-rest/database"
	"github.com/ivancl4udio/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
