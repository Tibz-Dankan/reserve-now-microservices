package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func updateRoomAmenity(w http.ResponseWriter, r *http.Request) {

}

func updateRoomAmenityRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/update-room/{id}", updateRoomAmenity).Methods("PATCH")
}
