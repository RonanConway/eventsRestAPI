package registrations

import (
	"net/http"
	"strconv"

	"github.com/RonanConway/eventsRestAPI/models"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register event"})
		return
	}

	// If all good, sending back sucess response.
	context.JSON(http.StatusCreated, gin.H{"message": "Registered event successfully!"})

}

func CancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to cancel registration"})
		return
	}

	// If all good, sending back sucess response.
	context.JSON(http.StatusOK, gin.H{"message": "Event Cancelled successfully!"})

}
