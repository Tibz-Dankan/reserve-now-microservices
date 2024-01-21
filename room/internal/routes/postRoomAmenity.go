package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/models"
	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/services"
	"github.com/gorilla/mux"
)

func postRoomAmenity(w http.ResponseWriter, r *http.Request) {
	amenity := models.Amenity{}

	fmt.Println("r.Body info from the frontend", r.Body) // To be removed
	err := json.NewDecoder(r.Body).Decode(&amenity)
	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}

	fmt.Println("amenity data", amenity)

	// user, ok := r.Context().Value("AuthUser").(middlewares.UserKey)
	// if !ok {
	// 	// If user details are not found in the context, handle accordingly
	// 	http.Error(w, "User details not found", http.StatusInternalServerError)
	// 	return
	// }
	// amenity.UpdatedByUserId = user.userId

	if amenity.Item == "" {
		services.AppError("Missing room name or type!", 400, w)
		return
	}

	savedAmenity, err := amenity.FindByItem(amenity.Item)
	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}

	if savedAmenity.ID > 0 {
		services.AppError(amenity.Item+" already added!", 400, w)
		return
	}

	fmt.Println("About to create a room")

	amenityId, err := amenity.Create(amenity)

	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}

	newAmenity := map[string]interface{}{
		"id":        amenityId,
		"item":      amenity.Item,
		"createdAt": amenity.CreatedAt,
		"updatedAt": amenity.UpdatedAt,
	}

	data := map[string]interface{}{"amenity": newAmenity}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Amenity created successfully!",
		"data":    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func postRoomAmenityRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/post-room-amenity", postRoomAmenity).Methods("POST")
}
