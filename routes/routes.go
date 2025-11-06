package routes

import (
	"github.com/RonanConway/eventsRestAPI/events"
	"github.com/RonanConway/eventsRestAPI/middlewares"
	"github.com/RonanConway/eventsRestAPI/registrations"
	"github.com/RonanConway/eventsRestAPI/users"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//Event routes
	server.GET("/events", events.GetEvents)
	server.GET("/events/:id", events.GetEvent)

	// This group will always run the authentication middleware
	// to ensure we have a valid JWT token therefore they are protected resources.
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", events.CreateEvent)
	authenticated.PUT("/events/:id", events.UpdateEvent)
	authenticated.DELETE("/events/:id", events.DeleteEvent)
	//event registrations endpoints
	authenticated.POST("/events/:id/register", registrations.RegisterForEvent)
	authenticated.DELETE("/events/:id/register", registrations.CancelRegistration)

	//User routes
	server.POST("/signup", users.Signup)
	server.POST("/login", users.Login)
	server.GET("/users", users.GetUsers)
}
