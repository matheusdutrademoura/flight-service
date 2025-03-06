# Flight Service

A RESTful service that aggregates flight information from multiple providers (SkyScanner, Google Flights, and CheapFlights), comparing prices and durations to find the best options. ⚠️ [Important note about provider implementations](#important-note-about-providers)

## Features

- Concurrent flight search across multiple providers [`service/flight_service.go`](service/flight.go)
- Request caching using singleflight pattern [`service/flight_service.go`](service/flight.go)
- JWT-based authentication [`handler/auth_handler.go`](handler/auth.go)
- Price and duration comparison [`service/flight_service.go`](service/flight.go)
- Docker support [`Dockerfile`](Dockerfile)

## Prerequisites

- Go 1.22 or higher
- Docker

## Installation

### Local Development

1. Clone the repository
```bash
git clone https://github.com/matheusdutrademoura/flight-service.git
cd flight-service
```

2. Install dependencies
```bash
go mod download
```

3. Run the tests
```bash
go test ./...
```

### Docker

1. Build the image
```bash
docker build -t flight-service .
```

2. Run the container
```bash
docker run -p 8080:8080 -e JWT_SECRET="supersecretkey" flight-service
```

Or using an environment file:
```bash
docker run --env-file .env -p 8080:8080 flight-service
```

## Configuration

### Environment Variable

- `JWT_SECRET`: Secret key for JWT token generation

## API Endpoints

### Authentication

#### Login
```http
POST /auth/login
Content-Type: application/json

{
    "username": "admin",
    "password": "admin123"
}
```

Default credentials:
- admin/admin123
- user/user123

Response:
```json
{
    "token": "your.jwt.token"
}
```

### Flights

#### Search Flights
```http
GET /flights/search?origin=NYC&destination=LAX&date=2024-01-01
Authorization: Bearer <your-jwt-token>
```

Response:
```json
{
    "cheapestFlight": {
        "provider": "CheapFlights",
        "price": 150.50,
        "duration": "2h0m",
        "departureTime": "2025-04-17T08:30:00Z"
    },
    "fastestFlight": {
        "provider": "GoogleFlights",
        "price": 220.25,
        "duration": "1h45m",
        "departureTime": "2025-04-17T16:20:00Z"
    },
    "providers": [
        // List of all flights from all providers
    ]
}
```
## Testing with curl

1. Get authentication token:
```bash
curl -X POST http://localhost:8080/auth/login 
  -H "Content-Type: application/json" 
  -d '{"username": "admin", "password": "admin123"}'
```

2. Search flights (combined command):
```bash
token=$(curl -s -X POST http://localhost:8080/auth/login 
  -H "Content-Type: application/json" 
  -d '{"username": "admin", "password": "admin123"}' | jq -r '.token') && 
curl "http://localhost:8080/flights/search?origin=NYC&destination=LAX&date=2024-01-01" 
  -H "Authorization: Bearer $token" | jq
```

3. Or use one single command to generate the token and call flights/search using jq:
```bash
curl -X POST http://localhost:8080/auth/login -H "Content-Type: application/json" -d '{"username": "admin", "password": "admin123"}' | jq -r '.token' | xargs -I {} curl "http://localhost:8080/flights/search?origin=NYC&destination=LAX&date=2024-01-01" -H "Authorization: Bearer {}" | jq
```

## Project Structure

```
flight-service/
├── handler/        # HTTP handlers
├── model/          # Data models
├── provider/       # Flight providers implementations
├── service/        # Business logic
├── main.go        # Application entry point
└── Dockerfile     # Docker configuration
```

## Important Note About Providers

This service uses mock implementations for all providers due to API availability constraints. I strongly suggest that you update your challenge request.

- **Google Flights** [`provider/google_flights.go`](provider/google_flights.go): The API was officially discontinued by Google on April 10, 2018, with no direct replacement available.
- **SkyScanner** [`provider/sky_scanner.go`](provider/sky_scanner.go): Requires a formal business partnership process:
  - Submit application to Skyscanner's Partnerships team
  - Go through business review process
  - Wait for team evaluation
  - Receive API credentials if approved
- **CheapFlights** [`provider/cheap_flights.go`](provider/cheap_flights.go): No information about a public API was found

All provider implementations in this project use mock data to demonstrate the service's architecture and functionality.

### Adding a New Provider

1. Implement the `FlightFetcher` interface in the provider package
2. Add the provider to the service initialization in `main.go`