package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Tibz-Dankan/reserve-now-microservices/auth/internal/models"
	"github.com/Tibz-Dankan/reserve-now-microservices/auth/internal/services"
	"github.com/gorilla/mux"
)

func signIn(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}

	password := user.Password

	user, err = user.FindByEMail(user.Email)
	if err != nil {
		services.AppError(err.Error(), 500, w)
		return
	}

	if user.ID == 0 {
		services.AppError("Email is not registered!", 400, w)
		return
	}

	passwordMatches, err := user.PasswordMatches(password)
	if err != nil {
		services.AppError(err.Error(), 500, w)
		return
	}

	if !passwordMatches {
		services.AppError("Invalid password!", 400, w)
		return
	}

	accessToken, err := services.SignJWTToken(user.ID)
	if err != nil {
		services.AppError(err.Error(), 500, w)
		return
	}

	userMap := map[string]interface{}{
		"id":      user.ID,
		"name":    user.Name,
		"email":   user.Email,
		"role":    user.Role,
		"country": user.Country,
	}
	response := map[string]interface{}{
		"status":      "success",
		"message":     "Sign in successfully",
		"accessToken": accessToken,
		"user":        userMap,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func SignInRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/auth/signin", signIn).Methods("POST")
}
