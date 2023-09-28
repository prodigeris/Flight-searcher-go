package main

import (
	"sort"
)

func generateItineraries(flights []Flight, maxDaysDifference int) []Itinerary {
	var itineraries []Itinerary

	homeAirports := map[string]bool{"VNO": true, "KUN": true}

	for _, departure := range flights {
		if _, ok := homeAirports[departure.FromAirport]; !ok {
			// trip doesn't start in the home airport
			continue
		}
		for _, returnFlight := range flights {
			if _, ok := homeAirports[returnFlight.ToAirport]; !ok {
				// trip doesn't end in the home airport
				continue
			}
			if departure.ToAirport != returnFlight.FromAirport {
				// trip is not connected
				continue
			}

			daysDifference := int(returnFlight.FlightDate.Sub(departure.FlightDate).Hours() / 24)
			if daysDifference > maxDaysDifference {
				// trip is too long
				continue
			}

			if departure.FlightDate.After(returnFlight.FlightDate) {
				// departure is after the return
				continue
			}

			itinerary := Itinerary{
				DepartureFlight: departure,
				ReturnFlight:    returnFlight,
				TotalPrice:      departure.Price + returnFlight.Price,
			}
			itineraries = append(itineraries, itinerary)
		}
	}

	return itineraries
}

func sortByPrice(its []Itinerary) {
	sort.Slice(its, func(i, j int) bool {
		return its[i].TotalPrice < its[j].TotalPrice
	})
}
