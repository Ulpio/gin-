package main

import (
	"github.com/Ulpio/gin-api-golang/database"
	"github.com/Ulpio/gin-api-golang/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequests()
}
