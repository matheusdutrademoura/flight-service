package provider

import (
	"flight-service/model"
	"time"
)

const CheapFlightsAPIURL = "to-be-defined-%s-%s-%s"

type CheapFlights struct {
	Name string
}

type CheapFlightsResponse struct {
	Data []struct {
		Price         float64       `json:"price"`
		Duration      time.Duration `json:"duration"`
		DepartureTime time.Time     `json:"departureTime"`
	} `json:"data"`
}

func NewCheapFlights() *CheapFlights {
	return &CheapFlights{
		Name: "CheapFlights",
	}
}

/*
I didn't find any API for cheapflights.com ¯\_(ツ)_/¯
I suggest you to update your challenge request
*/
func (c *CheapFlights) GetFlights(origin, destination string, date time.Time) ([]*model.Flight, error) {
	// Dummy implementation
	baseTime := time.Date(2025, time.March, 28, 0, 0, 0, 0, time.UTC)
	return []*model.Flight{
		{
			Provider:      c.Name,
			Price:         150.50,
			Duration:      2 * time.Hour,
			DepartureTime: baseTime.AddDate(0, 0, 2).Add(8*time.Hour + 30*time.Minute),
		},
		{
			Provider:      c.Name,
			Price:         200.75,
			Duration:      2*time.Hour + 30*time.Minute,
			DepartureTime: baseTime.AddDate(0, 0, 5).Add(14*time.Hour + 15*time.Minute),
		},
	}, nil
}
