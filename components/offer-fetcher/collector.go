package main

import (
	"database/sql"
	"errors"
	"github.com/airheartdev/duffel"
	"log"
	"os"
	"slices"
	"strconv"
)

func getOffer(search Search) (int, error) {
	//time.Sleep(1 * time.Second)
	offers := getOffers(duffel.New(os.Getenv("DUFFEL_TOKEN")), search.FromAirport, search.ToAirport, search.Date, 5)
	offerAmounts := make([]int, 0)
	for _, offer := range offers {
		amount, err := strconv.ParseFloat(offer.RawTotalAmount, 2)
		if err != nil {
			return 0, err
		}
		offerAmounts = append(offerAmounts, int(amount*100))
	}

	if len(offerAmounts) < 1 {
		return 0, errors.New("no offers found")
	}

	return slices.Min(offerAmounts), nil
}

func consumeSearch(db *sql.DB, search Search) error {
	log.Printf("Consuming search: %v", search)

	price, err := getOffer(search)
	if err != nil {
		return err
	}
	err = insertFlight(db, search.FromAirport, search.ToAirport, price, search.Date)
	if err != nil {
		return err
	}
	return nil
}
