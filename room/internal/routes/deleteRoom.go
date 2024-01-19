package routes

import (
	"net/http"
	"strconv"

	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/models"
	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/services"
	"github.com/gorilla/mux"
)

func deleteRoom(w http.ResponseWriter, r *http.Request) {
	// roomImage := models.RoomImage{}
	room := models.Room{}

	roomId := mux.Vars(r)["id"]

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
		services.AppError("Couldn't find image of provided roomImageId", 404, w)
		return
	}

	// TODO: first delete all room images from firebase storage
	// TODO: then delete room images from in db
	// TODO: finally delete room
	// FUTURE TODO: to delete room beds

}

func deleteRoomRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/delete-room/{id}", deleteRoom).Methods("DELETE")
}
