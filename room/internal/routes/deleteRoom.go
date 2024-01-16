package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func deleteRoom(w http.ResponseWriter, r *http.Request) {
	//   TODO: delete room operation here

}

func deleteRoomRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/delete-room", deleteRoom).Methods("DELETE")
}
