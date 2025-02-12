package main

import (
	"log"
	"net/http"
)

func main() {
	// Register HTTP handlers.
	http.HandleFunc("/items", itemsHandler)
	http.HandleFunc("/items/", itemHandler) // Handles endpoints like /items/{id}

	log.Println("Server starting on :8080")
	// Start the server.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
