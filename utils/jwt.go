package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

// The token thats generated here needs to be sent back in the response
// upon a successful login. Once logged in, the token needs to be passed
// in subsequent requests in order to access protected resources. If a
// protected resources does not find the JWT token in the header of the
// request then it should deny access to that resource.
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
