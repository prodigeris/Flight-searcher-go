package main

import (
	"testing"
	"time"
)

func TestGenerateItineraries(t *testing.T) {
	// Create some sample flight data
	flights := []Flight{
		{
			FromAirport: "VNO",
			ToAirport:   "RIX",
			Price:       100,
			FlightDate:  time.Now(),
		},
		{
			FromAirport: "RIX",
			ToAirport:   "VNO",
			Price:       110,
			FlightDate:  time.Now().AddDate(0, 0, 1).Add(time.Hour * 2), // 1 day difference
		},
		{
			FromAirport: "KUN",
			ToAirport:   "RIX",
			Price:       120,
			FlightDate:  time.Now(),
		},
		{
			FromAirport: "RIX",
			ToAirport:   "KUN",
			Price:       130,
			FlightDate:  time.Now().AddDate(0, 0, 3), // 3 days difference
		},
	}

	t.Run("Valid Itineraries", func(t *testing.T) {
		// Test valid itineraries with max days difference of 2
		maxDaysDifference := 2
		itineraries := generateItineraries(flights, maxDaysDifference)

		// Assert the number of valid itineraries
		if len(itineraries) != 2 {
			t.Errorf("Expected 2 valid itineraries, but got %d", len(itineraries))
		}

		// Add more specific assertions as needed for your test cases
	})

	t.Run("No Valid Itineraries", func(t *testing.T) {
		// Test when there are no valid itineraries
		maxDaysDifference := 0
		itineraries := generateItineraries(flights, maxDaysDifference)

		// Assert that there are no valid itineraries
		if len(itineraries) != 0 {
			t.Errorf("Expected 0 valid itineraries, but got %d", len(itineraries))
			t.Errorf("%s - %s", itineraries[0].DepartureFlight.FromAirport, itineraries[0].DepartureFlight.ToAirport)
			t.Errorf("%s - %s", itineraries[0].ReturnFlight.FromAirport, itineraries[0].ReturnFlight.ToAirport)
		}
	})

	// Add more test cases as needed
}

func TestSortByPrice(t *testing.T) {
	itineraries := []Itinerary{
		{
			TotalPrice: 200,
		},
		{
			TotalPrice: 100,
		},
		{
			TotalPrice: 400,
		},
		{
			TotalPrice: 222,
		},
	}

	// Call the function being tested
	sortByPrice(itineraries)

	// Assert that the itineraries are sorted by price in ascending order
	expectedPrices := []int{100, 200, 222, 400}
	for i, itinerary := range itineraries {
		if itinerary.TotalPrice != expectedPrices[i] {
			t.Errorf("Expected itinerary %d to have price %d, but got %d", i+1, expectedPrices[i], itinerary.DepartureFlight.Price)
		}
	}
}

func TestGenerateItinerariesHomeAirports(t *testing.T) {
	// Create some sample flight data
	flights := []Flight{
		{
			FromAirport: "VNO",
			ToAirport:   "RIX",
			Price:       100,
			FlightDate:  time.Now(),
		},
		{
			FromAirport: "RIX",
			ToAirport:   "VNO",
			Price:       110,
			FlightDate:  time.Now().AddDate(0, 0, 1),
		},
		{
			FromAirport: "KUN",
			ToAirport:   "RIX",
			Price:       120,
			FlightDate:  time.Now(),
		},
		{
			FromAirport: "RIX",
			ToAirport:   "KUN",
			Price:       130,
			FlightDate:  time.Now().AddDate(0, 0, 2),
		},
		{
			FromAirport: "VNO",
			ToAirport:   "RIX",
			Price:       140,
			FlightDate:  time.Now(),
		},
		{
			FromAirport: "RIX",
			ToAirport:   "KUN",
			Price:       150,
			FlightDate:  time.Now().AddDate(0, 0, 3),
		},
	}

	// Define the home airports
	homeAirports := map[string]bool{"VNO": true, "KUN": true}

	// Generate itineraries
	itineraries := generateItineraries(flights, 2)

	// Assert that all itineraries start at home airports
	for _, itinerary := range itineraries {
		departureAirport := itinerary.DepartureFlight.FromAirport

		if _, ok := homeAirports[departureAirport]; !ok {
			t.Errorf("Itinerary does not start at a home airport: %s", departureAirport)
		}
	}

	// Assert that all itineraries end at home airports
	for _, itinerary := range itineraries {
		returnAirport := itinerary.ReturnFlight.ToAirport

		if _, ok := homeAirports[returnAirport]; !ok {
			t.Errorf("Itinerary does not end at a home airport: %s", returnAirport)
		}
	}
}
