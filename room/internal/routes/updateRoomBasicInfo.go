package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func updateRoomBasicInfo(w http.ResponseWriter, r *http.Request) {
	//   TODO: update room basic information  operation here

}

func updateRoomBasicInfoRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/update-room", updateRoomBasicInfo).Methods("PATCH")
}
