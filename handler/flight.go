package handler

import (
	"encoding/json"
	"flight-service/model"
	"flight-service/provider"
	"flight-service/service"
	"net/http"
	"time"
)

func SearchFlightsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	origin := query.Get("origin")
	destination := query.Get("destination")
	dateStr := query.Get("date")

	if origin == "" || destination == "" || dateStr == "" {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	date, err := time.Parse(service.DateFormat, dateStr)
	if err != nil {
		http.Error(w, "Invalid date format. Please use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	searchParams := &model.SearchParams{
		Origin:      origin,
		Destination: destination,
		Date:        date,
	}

	fs := service.NewFlightService(
		provider.NewSkyScanner(),
		provider.NewGoogleFlights(),
		provider.NewCheapFlights(),
	)

	flights := fs.GetFlights(searchParams)
	comparison := fs.CompareFlights(flights)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comparison)
}
