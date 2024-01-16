package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func postRoomImage(w http.ResponseWriter, r *http.Request) {
	//   TODO: post room image operation here

}

func postRoomImageRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/post-room-image", postRoomImage).Methods("POST")
}
