package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Tibz-Dankan/reserve-now-microservices/internal/models"
	"github.com/Tibz-Dankan/reserve-now-microservices/internal/services"
	"github.com/gorilla/mux"
)

func forgotPassword(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}

	user, err = user.FindByEMail(user.Email)
	if err != nil {
		services.AppError(err.Error(), 500, w)
		return
	}

	if user.ID == 0 {
		services.AppError("We couldn't find user with provided email!", 400, w)
		return
	}

	// generate token (base64 string) use methods for the user struct
	// send it email service
	// send json payload {type: resetPassword, token, userName}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Reset token sent to mail",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func ForgotPasswordRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/auth/forgot-password", forgotPassword).Methods("POST")
}
