package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func unPublishRooms(w http.ResponseWriter, r *http.Request) {
	//   TODO: un publish room  operation here

}

func unPublishRoomsRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/search-rooms", unPublishRooms).Methods("PATCH")
}
