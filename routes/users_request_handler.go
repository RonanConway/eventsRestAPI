package routes

import (
	"net/http"

	"github.com/RonanConway/eventsRestAPI/models"
	"github.com/RonanConway/eventsRestAPI/utils"
	"github.com/gin-gonic/gin"
)

// User need to enter the email and password they signed up with or get rejected
// with invalid creds message. The password is hashed so a call out to Validate the
// credentials is required.
func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user request data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user"})
	}

	// Sending back the JWT token in the 200 response.
	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})

}

func signup(context *gin.Context) {
	var user models.User

	// user should be set by binding the incoming request
	// data to it like was done in the events handler
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse user request data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user to database!"})
		return
	}

	// Send back the sucess response
	context.JSON(http.StatusCreated, gin.H{"message": "User signup successful!"})

}

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch users, try again later."})
		return
	}
	context.JSON(http.StatusOK, users)
}
