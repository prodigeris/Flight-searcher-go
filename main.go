package main

import (
	"context"
	"fmt"
	"github.com/airheartdev/duffel"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {

	// Load connection string from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	dfl := duffel.New(os.Getenv("DUFFEL_TOKEN"))
	departureDate, _ := time.Parse(duffel.DateFormat, "2023-09-18")
	departureDate2, _ := time.Parse(duffel.DateFormat, "2023-09-25")
	maxConn := 0

	input := duffel.OfferRequestInput{
		Passengers: []duffel.OfferRequestPassenger{{Type: "adult"}},
		Slices: []duffel.OfferRequestSlice{
			{Origin: "VNO", Destination: "SFO", DepartureDate: duffel.Date(departureDate)},
			{Origin: "SFO", Destination: "KUN", DepartureDate: duffel.Date(departureDate2)},
		},
		CabinClass:      "economy",
		MaxConnections:  &maxConn,
		ReturnOffers:    true,
		SupplierTimeout: 10000,
	}

	request, err := dfl.CreateOfferRequest(context.Background(), input)
	if err != nil {
		fmt.Println(err)
	}

	for _, offer := range request.Offers {
		fmt.Println("Offer", offer.TotalAmount(), offer.RawTotalCurrency)
	}
}
