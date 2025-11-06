package events

import (
	"net/http"
	"strconv"

	"github.com/RonanConway/eventsRestAPI/models"
	"github.com/gin-gonic/gin"
)

var eventService EventsService = eventServiceImpl{}

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events, try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func GetEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrive event with Id"})
		return
	}

	context.JSON(http.StatusOK, event)

}

func CreateEvent(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	userId := context.GetInt64("userId")
	event.UserID = userId

	err = eventService.Save(&event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event, try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": event})
}

// retrieve the event from the database
// write the updated event back to the database.
func UpdateEvent(context *gin.Context) {
	//Get the event from the id passed in the request.
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id to update"})
		return
	}

	// Getting the matching event from the database
	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrive event with Id"})
		return
	}

	// Events should only be updated by users that created them.
	// The userId contained in the updated event must match the userId in the context for valid authed user who created the event.
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event!"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id to update"})
		return
	}

	updatedEvent.ID = eventId
	err = eventService.Update(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})
}

// Delete event by a DELETE request for a given ID.
func DeleteEvent(context *gin.Context) {
	//Get the event from the id passed in the request.
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id to delete"})
		return
	}

	userId := context.GetInt64("userId")

	// Getting the matching event from the database
	eventToDelete, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrive event with Id in order to delete it"})
		return
	}

	// Events should only be deleted by users that created them.
	// The userId contained in the event to be deleted must match the userId in the context for valid authed user who created the event.
	if eventToDelete.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event!"})
		return
	}

	err = eventService.Delete(eventToDelete)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted!"})

}
