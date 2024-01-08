package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/Tibz-Dankan/reserve-now-microservices/internal/services"
)

func resetPassword(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Password reset successfully"})
}

func ResetPasswordRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/auth/reset-password", resetPassword).Methods("PATCH")
}
