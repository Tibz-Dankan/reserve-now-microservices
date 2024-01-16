package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getAllRooms(w http.ResponseWriter, r *http.Request) {
	//   TODO: delete room operation here

}

func getAllRoomsRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/delete-room", getAllRooms).Methods("GET")
}
