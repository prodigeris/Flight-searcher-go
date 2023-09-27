package main

import (
	"database/sql"
	"fmt"
	"github.com/airheartdev/duffel"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
)

func collect(db *sql.DB, weekendCount int) {
	fmt.Println("Starting to consume inquiry for weekend count: ", weekendCount)
	dFlights, aFlights := GetAllDeparturesAndArrivals()
	friday, sunday := nextFridayAndSunday(time.Now())
	offerAmounts := make([]int, 0)
	for _, flight := range dFlights {
		fmt.Println("Starting to find offers for ", flight)
		offers := getOffers(duffel.New(os.Getenv("DUFFEL_TOKEN")), flight.FromAirport, flight.ToAirport, friday)
		for _, offer := range offers {
			amount, err := strconv.ParseFloat(offer.RawTotalAmount, 2)
			if err != nil {
				log.Fatalf("Failed to convert amount (%s)", offer.RawTotalAmount)
			}
			offerAmounts = append(offerAmounts, int(amount*100))
		}
		minOffer := slices.Min(offerAmounts)
		err := insertFlight(db, flight.FromAirport, flight.ToAirport, minOffer, friday)
		if err != nil {
			fmt.Println("Failed to insert flight")
			continue
		}
		break
	}
	fmt.Println(aFlights, sunday)

}
