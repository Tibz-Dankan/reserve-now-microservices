package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Tibz-Dankan/reserve-now-microservices/internal/routes"

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
	fmt.Println("Starting http server up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
