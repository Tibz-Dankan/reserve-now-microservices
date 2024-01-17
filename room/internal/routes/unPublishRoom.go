package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/models"
	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/services"
	"github.com/gorilla/mux"
)

func unPublishRooms(w http.ResponseWriter, r *http.Request) {
	room := models.Room{}
	roomId := mux.Vars(r)["id"]

	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}

	if roomId == "" {
		services.AppError("Please provide roomId!", 400, w)
		return
	}

	IntRoomId, err := strconv.Atoi(roomId)
	if err != nil {
		services.AppError(err.Error(), 500, w)
		return
	}

	savedRoom, err := room.FindOne(IntRoomId)
	if err != nil {
		services.AppError(err.Error(), 500, w)
		return
	}

	if savedRoom.ID == 0 {
		services.AppError("Room of supplied id doesn't exist!", 404, w)
		return
	}

	if !room.IsPublished() {
		services.AppError("Room already unpublished!", 400, w)
		return
	}

	err = room.UpdateAsUnPublished(IntRoomId)
	if err != nil {
		services.AppError(err.Error(), 500, w)
		return
	}

	room, err = room.FindOne(IntRoomId)
	if err != nil {
		services.AppError(err.Error(), 500, w)
		return
	}

	data := map[string]interface{}{"room": room}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Room unpublished successfully!",
		"data":    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func unPublishRoomsRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/unpublish-room/{id}", unPublishRooms).Methods("PATCH")
}
