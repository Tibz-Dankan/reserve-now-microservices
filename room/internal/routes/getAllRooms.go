package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/models"
	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/services"
	"github.com/gorilla/mux"
)

func getAllRooms(w http.ResponseWriter, r *http.Request) {
	room := models.Room{}

	rooms, err := room.FindAll()
	if err != nil {
		services.AppError(err.Error(), 500, w)
	}

	data := map[string]interface{}{"rooms": rooms}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Rooms fetched successfully!",
		"data":    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func getAllRoomsRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/get-all-rooms", getAllRooms).Methods("GET")
}
