package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func deleteRoomAmenity(w http.ResponseWriter, r *http.Request) {

	amenityId := mux.Vars(r)["id"]
	fmt.Println("AmenityIs", amenityId)
	// Delete amend operation here

}

func deleteRoomAmenityRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/delete-room-amenity/{id}", deleteRoomAmenity).Methods("DELETE")
}
