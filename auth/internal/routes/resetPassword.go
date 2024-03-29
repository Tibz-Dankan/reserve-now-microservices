package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Tibz-Dankan/reserve-now-microservices/auth/internal/models"
	"github.com/Tibz-Dankan/reserve-now-microservices/auth/internal/services"
	"github.com/gorilla/mux"
)

func resetPassword(w http.ResponseWriter, r *http.Request) {

	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}

	newPassword := user.Password
	token := mux.Vars(r)["resetToken"]

	user, err = user.FindByPasswordResetToken(token)
	if err != nil {
		services.AppError(err.Error(), 500, w)
		return
	}

	if user.ID == 0 {
		services.AppError("Invalid or expired reset token!", 400, w)
		return
	}

	user.ResetPassword(newPassword)

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
		"message":     "Password reset successfully",
		"accessToken": accessToken,
		"user":        userMap,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func ResetPasswordRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/auth/reset-password/{resetToken}", resetPassword).Methods("PATCH")
}
