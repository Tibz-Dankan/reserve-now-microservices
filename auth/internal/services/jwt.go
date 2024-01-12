package services

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func SignJWTToken(userId int) *jwt.Token {
	// Create a new token object, specifying signing method and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"nbf":    time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	return token
}
