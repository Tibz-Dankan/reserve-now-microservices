package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func searchRooms(w http.ResponseWriter, r *http.Request) {
	//   TODO: search rooms  operation here

}

func searchRoomsRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/search-rooms", searchRooms).Methods("GET")
}
