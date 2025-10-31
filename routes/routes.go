package routes

import (
	"github.com/RonanConway/eventsRestAPI/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//Event routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// This group will always run the authentication middleware
	// to ensure we have a valid JWT token therefore they are protected resources.
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	//User routes
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/users", getUsers)
}
