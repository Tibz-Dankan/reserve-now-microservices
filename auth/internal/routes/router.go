package routes

import (
	"github.com/Tibz-Dankan/reserve-now-microservices/internal/middlewares"

	"github.com/gorilla/mux"
)

func AppRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(middlewares.Logger)
	// router.Use(middlewares.Auth)

	SignUpRoute(router)
	SignInRoute(router)
	ForgotPasswordRoute(router)
	ResetPasswordRoute(router)

	return router
}
