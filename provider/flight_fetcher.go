package provider

import (
	"flight-service/model"
	"time"
)

type FlightFetcher interface {
	GetFlights(origin, destination string, date time.Time) ([]*model.Flight, error)
}
