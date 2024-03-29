package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/models"
	"github.com/Tibz-Dankan/reserve-now-microservices/room/internal/routes"

	"github.com/rs/cors"
)

func main() {
	router := routes.AppRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "production_url"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete},
		AllowCredentials: true,
		Debug:            true,
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(router)

	http.Handle("/", handler)

	models.DBAutoMigrate()

	fmt.Println("room server running  on 8001")
	log.Fatal(http.ListenAndServe(":8001", nil))
}
