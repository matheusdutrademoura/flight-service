package service

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"flight-service/model"
	"flight-service/provider"
)

func TestFlightServiceImpl_CompareFlights(t *testing.T) {
	baseTime := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		flights  []*model.Flight
		expected *model.FlightComparison
	}{
		{
			name:    "empty flights list",
			flights: []*model.Flight{},
			expected: &model.FlightComparison{
				CheapestFlight: nil,
				FastestFlight:  nil,
				Providers:      nil,
			},
		},
		{
			name: "single flight",
			flights: []*model.Flight{
				{
					Provider:      "Provider1",
					Price:         100.0,
					Duration:      2 * time.Hour,
					DepartureTime: baseTime,
				},
			},
			expected: &model.FlightComparison{
				CheapestFlight: &model.Flight{
					Provider:      "Provider1",
					Price:         100.0,
					Duration:      2 * time.Hour,
					DepartureTime: baseTime,
				},
				FastestFlight: &model.Flight{
					Provider:      "Provider1",
					Price:         100.0,
					Duration:      2 * time.Hour,
					DepartureTime: baseTime,
				},
				Providers: []*model.Flight{
					{
						Provider:      "Provider1",
						Price:         100.0,
						Duration:      2 * time.Hour,
						DepartureTime: baseTime,
					},
				},
			},
		},
		{
			name: "multiple flights with different prices and durations",
			flights: []*model.Flight{
				{
					Provider:      "Provider1",
					Price:         200.0,
					Duration:      3 * time.Hour,
					DepartureTime: baseTime,
				},
				{
					Provider:      "Provider2",
					Price:         150.0,
					Duration:      4 * time.Hour,
					DepartureTime: baseTime.Add(1 * time.Hour),
				},
				{
					Provider:      "Provider3",
					Price:         300.0,
					Duration:      2 * time.Hour,
					DepartureTime: baseTime.Add(2 * time.Hour),
				},
			},
			expected: &model.FlightComparison{
				CheapestFlight: &model.Flight{
					Provider:      "Provider2",
					Price:         150.0,
					Duration:      4 * time.Hour,
					DepartureTime: baseTime.Add(1 * time.Hour),
				},
				FastestFlight: &model.Flight{
					Provider:      "Provider3",
					Price:         300.0,
					Duration:      2 * time.Hour,
					DepartureTime: baseTime.Add(2 * time.Hour),
				},
				Providers: []*model.Flight{
					{
						Provider:      "Provider1",
						Price:         200.0,
						Duration:      3 * time.Hour,
						DepartureTime: baseTime,
					},
					{
						Provider:      "Provider2",
						Price:         150.0,
						Duration:      4 * time.Hour,
						DepartureTime: baseTime.Add(1 * time.Hour),
					},
					{
						Provider:      "Provider3",
						Price:         300.0,
						Duration:      2 * time.Hour,
						DepartureTime: baseTime.Add(2 * time.Hour),
					},
				},
			},
		},
	}

	service := &FlightServiceImpl{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.CompareFlights(tt.flights)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFlightServiceImpl_GetFlights_Concurrent(t *testing.T) {
	s := NewFlightService(
		provider.NewSkyScanner(),
		provider.NewGoogleFlights(),
		provider.NewCheapFlights(),
	)
	params := &model.SearchParams{
		Origin:      "NYC",
		Destination: "LAX",
		Date:        time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	const numConcurrentCalls = 10
	var wg sync.WaitGroup
	var results [][]*model.Flight
	var resultsMutex sync.Mutex

	// Make concurrent calls
	wg.Add(numConcurrentCalls)
	for i := 0; i < numConcurrentCalls; i++ {
		go func() {
			defer wg.Done()
			flights := s.GetFlights(params)

			resultsMutex.Lock()
			results = append(results, flights)
			resultsMutex.Unlock()
		}()
	}

	// Wait with timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// Success path
	case <-time.After(5 * time.Second):
		t.Fatal("Test timed out")
	}

	// Verify results
	assert.Equal(t, numConcurrentCalls, len(results), "Should have results from all calls")

	// All results should be identical due to singleflight
	for i := 1; i < len(results); i++ {
		assert.Equal(t, results[0], results[i], "All results should be identical due to caching")
	}

	// Verify content of flights
	firstResult := results[0]
	assert.NotEmpty(t, firstResult, "Should have flights")

	providers := make(map[string]bool)
	for _, flight := range firstResult {
		providers[flight.Provider] = true
	}
	assert.Equal(t, 3, len(providers), "Should have flights from all three providers")
}
