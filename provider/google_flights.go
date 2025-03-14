package provider

import (
	"flight-service/model"
	"time"
)

const GoogleFlightsAPIURL = "to-be-defined-%s-%s-%s"

type GoogleFlights struct {
	Name string
}

type GoogleFlightsResponse struct {
	Data []struct {
		Price         float64       `json:"price"`
		Duration      time.Duration `json:"duration"`
		DepartureTime time.Time     `json:"departureTime"`
	} `json:"data"`
}

func NewGoogleFlights() *GoogleFlights {
	return &GoogleFlights{
		Name: "GoogleFlights",
	}
}

/*
Google ended access to the public-facing API in 2018. I suggest you to update your challenge request
*/
func (g *GoogleFlights) GetFlights(origin, destination string, date time.Time) ([]*model.Flight, error) {
	// Dummy implementation
	baseTime := time.Date(2025, time.April, 10, 0, 0, 0, 0, time.UTC)
	return []*model.Flight{
		{
			Provider:      g.Name,
			Price:         180.00,
			Duration:      2*time.Hour + 15*time.Minute,
			DepartureTime: baseTime.AddDate(0, 0, 3).Add(11*time.Hour + 45*time.Minute),
		},
		{
			Provider:      g.Name,
			Price:         220.25,
			Duration:      1*time.Hour + 45*time.Minute,
			DepartureTime: baseTime.AddDate(0, 0, 7).Add(16*time.Hour + 20*time.Minute),
		},
	}, nil
}
