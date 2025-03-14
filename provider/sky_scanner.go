package provider

import (
	"flight-service/model"
	"time"
)

const SkyScannerAPIURL = "to-be-defined-%s-%s-%s"

type SkyScanner struct {
	Name string
}

type SkyScannerResponse struct {
	Data []struct {
		Price         float64       `json:"price"`
		Duration      time.Duration `json:"duration"`
		DepartureTime time.Time     `json:"departureTime"`
	} `json:"data"`
}

func NewSkyScanner() *SkyScanner {
	return &SkyScanner{
		Name: "SkyScanner",
	}
}

/*
To get a Skyscanner API key, I need to:
- Submit an application to Skyscanner's Partnerships team
- Wait for the team to review my application
- If the application is successful, Skyscanner will contact me, in theory.

I suggest you to update your challenge request
*/
func (s *SkyScanner) GetFlights(origin, destination string, date time.Time) ([]*model.Flight, error) {
	// Dummy implementation
	baseTime := time.Date(2025, time.April, 5, 0, 0, 0, 0, time.UTC)
	return []*model.Flight{
		{
			Provider:      s.Name,
			Price:         165.75,
			Duration:      2*time.Hour + 30*time.Minute,
			DepartureTime: baseTime.AddDate(0, 0, 4).Add(7*time.Hour + 15*time.Minute),
		},
		{
			Provider:      s.Name,
			Price:         190.25,
			Duration:      2 * time.Hour,
			DepartureTime: baseTime.AddDate(0, 0, 8).Add(13*time.Hour + 40*time.Minute),
		},
	}, nil
}
