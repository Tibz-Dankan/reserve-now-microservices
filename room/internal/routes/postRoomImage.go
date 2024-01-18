package routes

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/services"
	"github.com/gorilla/mux"
)

func postRoomImage(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
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

	fmt.Println("imageUrl ===> ", imageUrl)
	// save image in the database

}

func postRoomImageRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/room/post-room-image", postRoomImage).Methods("POST")
}
