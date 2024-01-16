package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Tibz-Dankan/reserve-now-microservices/auth/internal/models"
	"github.com/Tibz-Dankan/reserve-now-microservices/auth/internal/services"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

func verifyAuth(w http.ResponseWriter, r *http.Request) {

	authorizationHeader := r.Header.Get("Authorization")

	if authorizationHeader == "" {
		services.AppError("Missing Authorization header", 401, w)
		return
	}
	headerParts := strings.SplitN(authorizationHeader, " ", 2)

	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		services.AppError("Invalid Authorization header", 401, w)
		return
	}
	bearerToken := headerParts[1]
	secretKey := os.Getenv("JWT_SECRET")
	var jwtSecretKey = []byte(secretKey)

	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			services.AppError("Unexpected signing method", 403, w)
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		services.AppError(err.Error(), 403, w)
		return
	}

	var userId int
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIdClaim, _ := claims["userId"].(float64)
		userId = int(userIdClaim)
	} else {
		services.AppError("Invalid Token, please login again", 403, w)
		return
	}
	fmt.Println("userId", userId)

	User := models.User{}

	user, err := User.FindOne(userId)
	if err != nil {
		services.AppError(err.Error(), 500, w)
		return
	}
	log.Println("authorized user", user)

	response := map[string]interface{}{
		"status":  "success",
		"message": "user is authorized",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func verifyAuthRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/auth/verify", verifyAuth).Methods("POST")
}
