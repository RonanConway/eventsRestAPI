package main

import (
	"github.com/RonanConway/eventsRestAPI/db"
	"github.com/RonanConway/eventsRestAPI/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost:8080

	//TODO, finish client
	// Client code
	// client := client.NewClient("http://localhost:8080")

}
