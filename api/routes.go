package api

import (
	"golang-blog-api/internal/handlers"
	"golang-blog-api/internal/services"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	router := mux.NewRouter()

	authService := &services.AuthServiceImpl{}
	createUserHandler := handlers.NewUserHandler(authService)

	router.HandleFunc("/api/register", createUserHandler.ServeHTTP).Methods("POST")
	// router.HandleFunc("/api/login", handlers.LoginUserHandler).Methods("POST")

	return router
}
