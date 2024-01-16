package routes

import (
	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/middlewares"

	"github.com/gorilla/mux"
)

func AppRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(middlewares.Logger)
	router.Use(middlewares.Auth)

	postRoomBasicInfoRoute(router)
	updateRoomBasicInfoRoute(router)
	postRoomImageRoute(router)
	updateRoomImageRoute(router)
	publishRoomRoute(router)
	unPublishRoomsRoute(router)
	getAllRoomsRoute(router)
	searchRoomsRoute(router)
	searchRoomsRoute(router)
	deleteRoomRoute(router)

	return router
}
