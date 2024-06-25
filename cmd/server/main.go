package main

import (
	"fmt"
	"log"
	"net/http"

	"medBackend/internal/handlers"
)

func main() {
	// HTTP mux yaratish
	mux := http.NewServeMux()

	// Routes qo'shish
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/about", handlers.AboutHandler)
	mux.HandleFunc("/register", handlers.RegisterHandler)
	mux.HandleFunc("/auth", handlers.AuthHandler)

	// Serverni tuzish va ishga tushirish
	port := ":8080"
	fmt.Printf("Server %s portida ishga tushirildi...\n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
