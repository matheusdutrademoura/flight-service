package handler

import (
	"encoding/json"
	"flight-service/model"
	"flight-service/service"
	"net/http"
	"time"
)

type FlightHandler struct {
	flightService service.FlightService
}

func NewFlightHandler(fs service.FlightService) *FlightHandler {
	return &FlightHandler{
		flightService: fs,
	}
}

func (h *FlightHandler) SearchFlights(w http.ResponseWriter, r *http.Request) {
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

	flights := h.flightService.GetFlights(searchParams)
	comparison := h.flightService.CompareFlights(flights)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comparison)
}
