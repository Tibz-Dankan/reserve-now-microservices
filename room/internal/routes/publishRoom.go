package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func publishRoom(w http.ResponseWriter, r *http.Request) {
	//   TODO: publish room  operation here

}

func publishRoomRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/publish-room", publishRoom).Methods("PATCH")
}
