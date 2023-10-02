package main

import (
	"database/sql"
	"errors"
	"github.com/airheartdev/duffel"
	"log"
	"os"
	"strconv"
)

func getOffer(search Search) (int, error) {
	offers := getOffers(duffel.New(os.Getenv("DUFFEL_TOKEN")), search.FromAirport, search.ToAirport, search.Date, 5)
	minOffer := 999999999
	for _, offer := range offers {
		amount, err := strconv.ParseFloat(offer.RawTotalAmount, 2)
		if err != nil {
			return 0, err
		}
		if int(amount*100) < minOffer {
			minOffer = int(amount * 100)
		}
	}

	if minOffer == 999999999 {
		return 0, errors.New("no offers found")
	}

	return minOffer, nil
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
