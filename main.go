package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Check and read environment variables
	testId := os.Getenv("testId")
	if testId == "" {
		fmt.Println("Environment variable 'testId' not set")
	}

	testPassword := os.Getenv("testPassword")
	if testPassword == "" {
		fmt.Println("Environment variable 'testPassword' not set")
	}

	// Set up HTTP server with separate endpoints for each environment variable
	http.HandleFunc("/testId", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"testId": testId})
	})

	http.HandleFunc("/testPassword", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"testPassword": testPassword})
	})

	// Start the server on port 9090
	log.Println("Server starting on port 9090...")
	log.Fatal(http.ListenAndServe(":9090", nil))
}
