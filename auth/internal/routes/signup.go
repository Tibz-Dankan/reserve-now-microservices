package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Tibz-Dankan/reserve-now-microservices/internal/models"
	"github.com/Tibz-Dankan/reserve-now-microservices/internal/services"
	"github.com/gorilla/mux"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}

	savedUser, err := user.FindByEMail(user.Email)
	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}

	fmt.Println("savedUser", savedUser)

	if savedUser.ID > 0 {
		services.AppError("Email already registered!", 400, w)
		return
	}

	// var userId int
	// userId, err := user.Create(user)
	_, err = user.Create(user)
	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}
	// TODO: login user upon creating account

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Signup successfully"})
}

func SignUpRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/auth/signup", signUp).Methods("POST")
}
