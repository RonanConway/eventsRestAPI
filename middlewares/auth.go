package middlewares

import (
	"net/http"

	"github.com/RonanConway/eventsRestAPI/utils"
	"github.com/gin-gonic/gin"
)

// This is executed in the middle of a request. If something goes wrong
// here we must stop so no other code on the server runs.
func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		// Calling this abort wilkl stop the chain.
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	userId, tokenErr := utils.VerifyToken(token)
	if tokenErr != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	// Setting this on the context here so it can be extracted in other handlers where its needed.
	context.Set("userId", userId)
	// This ensures that the next request handler in line will execute correctly.
	context.Next()

}
