package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/models"
	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/services"
	"github.com/gorilla/mux"
)

type CreateRoom struct {
	RoomName  string `json:"roomName"`
	RoomType  string `json:"roomType"`
	Capacity  string `json:"capacity"`
	Price     string `json:"price"`
	Amenities string `json:"amenities"`
	View      string `json:"view"`
}

func postRoomBasicInfo(w http.ResponseWriter, r *http.Request) {
	room := models.Room{}
	roomInputData := CreateRoom{}

	fmt.Println("r.Body info from the frontend", r.Body)
	// err := json.NewDecoder(r.Body).Decode(&room)
	err := json.NewDecoder(r.Body).Decode(&roomInputData)
	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}

	fmt.Println("room info from the frontend", room)

	if roomInputData.RoomName == "" || roomInputData.RoomType == "" {
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

	fmt.Println("About to create a room")

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
