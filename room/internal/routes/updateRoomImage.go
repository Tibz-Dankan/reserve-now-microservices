package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func updateRoomImage(w http.ResponseWriter, r *http.Request) {
	//   TODO: update room basic information  operation here

}

func updateRoomImageRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/update-room-image", updateRoomImage).Methods("PATCH")
}
