package main

import (
	"fmt"
	"github.com/airheartdev/duffel"
	"os"
	"time"
)

func collect(weekendCount int) {
	fmt.Println("Starting to consume inquiry for weekend count: ", weekendCount)
	dFlights, aFlights := GetAllDeparturesAndArrivals()
	friday, sunday := nextFridayAndSunday(time.Now())
	for _, flight := range dFlights {
		fmt.Println("Starting to find offers for ", flight)
		offers := getOffers(duffel.New(os.Getenv("DUFFEL_TOKEN")), flight.FromAirport, flight.ToAirport, friday)
		for _, offer := range offers {
			fmt.Println("Found offer for", offer.RawTotalAmount, offer.RawTotalCurrency)
		}
		break
	}
	fmt.Println(aFlights, sunday)

}
