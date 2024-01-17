package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/models"
	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/services"
	"github.com/gorilla/mux"
)

func updateRoomBasicInfo(w http.ResponseWriter, r *http.Request) {
	room := models.Room{}
	roomId := mux.Vars(r)["id"]

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
		services.AppError("Please provide at least the number of adults!", 400, w)
		return
	}

	if !room.IsValidRoomPrice() {
		services.AppError("Please provide the price amount and currency!", 400, w)
		return
	}

	IntRoomId, err := strconv.Atoi(roomId)
	if err != nil {
		services.AppError(err.Error(), 500, w)
		return
	}

	savedRoom, err := room.FindOne(IntRoomId)
	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}

	if savedRoom.ID == 0 {
		services.AppError("Room of supplied id doesn't exist!", 404, w)
		return
	}

	if savedRoom.RoomName != room.RoomName {
		savedRoom, err := room.FindByName(room.RoomName)

		if err != nil {
			services.AppError(err.Error(), 400, w)
			return
		}
		if savedRoom.ID > 0 {
			services.AppError("Can't Update to already existing room name!", 400, w)
			return
		}
	}

	err = room.Update(IntRoomId)
	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}

	updatedRoom := map[string]interface{}{
		"id":        roomId,
		"roomName":  room.RoomName,
		"roomType":  room.RoomType,
		"capacity":  room.Capacity,
		"price":     room.Price,
		"createdAt": room.CreatedAt,
	}

	data := map[string]interface{}{"room": updatedRoom}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Room updated successfully!",
		"data":    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updateRoomBasicInfoRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/update-room/{id}", updateRoomBasicInfo).Methods("PATCH")
}
