package routes

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/models"
	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/services"
	"github.com/gorilla/mux"
)

func postRoomImage(w http.ResponseWriter, r *http.Request) {
	roomImage := models.RoomImage{}
	roomId := mux.Vars(r)["id"]

	err := json.NewDecoder(r.Body).Decode(&roomImage)
	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}

	if roomId == "" {
		services.AppError("Please provide roomId!", 400, w)
		return
	}

	if roomImage.ViewType == "" {
		services.AppError("Please provide room view type!", 400, w)
		return
	}

	IntRoomId, err := strconv.Atoi(roomId)
	if err != nil {
		services.AppError(err.Error(), 500, w)
		return
	}

	err = r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		services.AppError("Unable to parse form", 400, w)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		services.AppError(err.Error(), 400, w)
		return
	}
	defer file.Close()

	randNumStr := strconv.Itoa(rand.Intn(9000) + 1000)
	filePath := "go/rooms/" + randNumStr + "_" + fileHeader.Filename

	upload := services.Upload{FilePath: filePath}

	imageUrl, err := upload.Add(file, fileHeader)
	if err != nil {
		services.AppError(err.Error(), 500, w)
		return
	}

	fmt.Println("imageUrl :", imageUrl)

	roomImage.RoomID = IntRoomId
	roomImage.URL = imageUrl
	roomImage.Path = filePath

	roomImageID, err := roomImage.Create(roomImage)
	if err != nil {
		services.AppError(err.Error(), 500, w)
		return
	}

	newRoomImage := map[string]interface{}{
		"id":        roomImageID,
		"roomId":    roomImage.RoomID,
		"url":       roomImage.URL,
		"viewType":  roomImage.ViewType,
		"createdAt": roomImage.CreatedAt,
	}

	data := map[string]interface{}{"roomImage": newRoomImage}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Room Image uploaded successfully!",
		"data":    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

func postRoomImageRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/post-room-image/{id}", postRoomImage).Methods("POST")
}
