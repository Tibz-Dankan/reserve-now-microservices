package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func postRoomBasicInfo(w http.ResponseWriter, r *http.Request) {
	//   TODO: post room operation here

}

func postRoomBasicInfoRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/post-room", postRoomBasicInfo).Methods("POST")
}
