package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/models"
	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/services"
	"github.com/gorilla/mux"
)

func postRoomBasicInfo(w http.ResponseWriter, r *http.Request) {
	room := models.Room{}

	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}

	if room.RoomName == "" || room.RoomType == "" {
		services.AppError("Missing room name or type!", 400, w)
		return
	}

	if !room.IsValidRoomCapacity() {
		services.AppError("Please provide at least the number of adults", 400, w)
		return
	}

	if !room.IsValidRoomPrice() {
		services.AppError("Please provide the price amount and currency", 400, w)
		return
	}

	savedRoom, err := room.FindByName(room.RoomName)
	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}

	if savedRoom.ID > 0 {
		services.AppError(room.RoomName+" already exists!", 400, w)
		return
	}

	roomId, err := room.Create(room)

	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}

	newRoom := map[string]interface{}{
		"id":        roomId,
		"roomName":  room.RoomName,
		"roomType":  room.RoomType,
		"capacity":  room.Capacity,
		"price":     room.Price,
		"createdAt": room.CreatedAt,
	}

	data := map[string]interface{}{"room": newRoom}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Room created successfully!",
		"data":    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func postRoomBasicInfoRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/post-room", postRoomBasicInfo).Methods("POST")
}
