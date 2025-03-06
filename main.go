package main

import (
	"flight-service/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/flights/search", handler.AuthMiddleware(handler.SearchFlightsHandler))
	http.HandleFunc("/auth/login", handler.LoginHandler)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
