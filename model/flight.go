package model

import (
	"encoding/json"
	"time"
)

type Flight struct {
	Provider          string        `json:"provider"`
	Price             float64       `json:"price"`
	Duration          time.Duration `json:"-"` // Hide the original duration
	DepartureTime     time.Time     `json:"departureTime"`
	FormattedDuration string        `json:"duration"` // Add formatted duration
}

// MarshalJSON implements custom JSON marshaling
func (f *Flight) MarshalJSON() ([]byte, error) {
	type Alias Flight
	return json.Marshal(&struct {
		*Alias
		Duration string `json:"duration"`
	}{
		Alias:    (*Alias)(f),
		Duration: f.Duration.String(),
	})
}

type FlightComparison struct {
	CheapestFlight *Flight   `json:"cheapestFlight"`
	FastestFlight  *Flight   `json:"fastestFlight"`
	Providers      []*Flight `json:"providers"`
}

type SearchParams struct {
	Origin      string    `form:"origin" binding:"required"`
	Destination string    `form:"destination" binding:"required"`
	Date        time.Time `form:"date" binding:"required,datetime=2006-01-02"`
}
