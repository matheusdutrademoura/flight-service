package provider

import (
	"flight-service/model"
)

type FlightFetcher interface {
	GetFlights(origin, destination, date string) ([]*model.Flight, error)
}
