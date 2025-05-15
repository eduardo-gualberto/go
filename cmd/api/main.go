package main

import (
	"fmt"
	"net/http"

	"github.com/eduardo-gualberto/go.git/gateways/handlers"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading env: %v", err)
	}
}

func main() {
	http.HandleFunc("GET /webhook", handlers.HandleAuthenticate)
	http.HandleFunc("POST /webhook", handlers.HandleMessage)
	http.ListenAndServe(":8080", nil)
}
