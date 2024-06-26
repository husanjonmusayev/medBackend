package main

import (
	"fmt"
	"log"
	"net/http"
	"your_module_name/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Create a new HTTP router using mux
	router := mux.NewRouter()

	// Define your routes
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/about", handlers.AboutHandler).Methods("GET")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/auth", handlers.AuthHandler).Methods("POST")

	// CORS configuration
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5500"}, // Add your frontend URL here
		AllowedMethods: []string{"GET", "POST"},           // Allow only GET and POST methods
	}).Handler(router)

	// Set up server
	port := ":8080"
	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, corsHandler))
}
