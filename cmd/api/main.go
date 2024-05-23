package api

import (
	"fmt"
	"log"
	"net/http"

	"kissandeat/cmd/handlers"
	"kissandeat/config"
	"kissandeat/internal/repository"
)

// StartServer starts the API server
func StartServer(dbConfig *config.Config, repo *repository.Repository) error {
	// Create router and handlers
	router := http.NewServeMux()
	authHandler := handlers.NewAuthHandler(repo)
	userHandler := handlers.NewUserHandler(repo)
	dishHandler := handlers.NewDishHandler(repo)

	// Define routes and handlers
	router.HandleFunc("/auth/login", authHandler.Login)
	router.HandleFunc("/auth/register", authHandler.Register)
	router.HandleFunc("/users", userHandler.ListUsers)
	router.HandleFunc("/users/{id}", userHandler.GetUser)
	router.HandleFunc("/dishes", dishHandler.ListDishes)
	router.HandleFunc("/dishes/{id}", dishHandler.GetDish)

	// Start the HTTP server
	log.Printf("Starting log API server on port %d", dbConfig.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", dbConfig.Port), router)
}
