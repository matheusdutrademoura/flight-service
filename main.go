package main

import (
	"flight-service/handler"
	"flight-service/provider"
	"flight-service/service"
	"log"
	"net/http"
)

func main() {
	flightService := service.NewFlightService(
		provider.NewSkyScanner(),
		provider.NewGoogleFlights(),
		provider.NewCheapFlights(),
	)

	flightHandler := handler.NewFlightHandler(flightService)

	http.HandleFunc("/auth/login", handler.LoginHandler)
	http.HandleFunc("/flights/search", handler.AuthMiddleware(flightHandler.SearchFlights))
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
