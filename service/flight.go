package service

import (
	"flight-service/model"
	"flight-service/provider"
	"fmt"
	"sync"

	"golang.org/x/sync/singleflight"
)

const (
	DateFormat = "2006-01-02"
)

type FlightService interface {
	GetFlights(params *model.SearchParams) []*model.Flight
	CompareFlights(flights []*model.Flight) *model.FlightComparison
}

type FlightServiceImpl struct {
	providers   []provider.FlightFetcher
	flightGroup singleflight.Group
}

func NewFlightService(providers ...provider.FlightFetcher) *FlightServiceImpl {
	return &FlightServiceImpl{
		providers:   providers,
		flightGroup: singleflight.Group{},
	}
}

func (s *FlightServiceImpl) GetFlights(params *model.SearchParams) []*model.Flight {
	key := fmt.Sprintf("%s-%s-%s", params.Origin, params.Destination, params.Date.Format(DateFormat))

	results, _, _ := s.flightGroup.Do(key, func() (interface{}, error) {
		return s.getFlights(params), nil
	})

	return results.([]*model.Flight)
}

func (s *FlightServiceImpl) getFlights(params *model.SearchParams) []*model.Flight {
	var wg sync.WaitGroup
	results := make(chan *model.Flight)
	dateStr := params.Date.Format(DateFormat)

	fetcherFunc := func(p provider.FlightFetcher) {
		defer wg.Done()
		flights, err := p.GetFlights(params.Origin, params.Destination, dateStr)
		if err != nil {
			return
		}
		for _, flight := range flights {
			results <- flight
		}
	}

	wg.Add(len(s.providers))
	for _, provider := range s.providers {
		go fetcherFunc(provider)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var flights []*model.Flight
	for flight := range results {
		flights = append(flights, flight)
	}
	return flights
}

func (s *FlightServiceImpl) CompareFlights(flights []*model.Flight) *model.FlightComparison {
	if len(flights) == 0 {
		return &model.FlightComparison{}
	}

	cheapest := flights[0]
	fastest := flights[0]

	for _, flight := range flights[1:] {
		if flight.Price < cheapest.Price {
			cheapest = flight
		}
		if flight.Duration < fastest.Duration {
			fastest = flight
		}
	}

	return &model.FlightComparison{
		CheapestFlight: cheapest,
		FastestFlight:  fastest,
		Providers:      flights,
	}
}
