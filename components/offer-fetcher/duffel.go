package main

import (
	"context"
	"fmt"
	"github.com/airheartdev/duffel"
	"time"
)

type duffelClient interface {
	CreateOfferRequest(ctx context.Context, input duffel.OfferRequestInput) (*duffel.OfferRequest, error)
}

func getOffers(dfl duffelClient, origin string, destination string, departureDate time.Time, maxRetries int) []duffel.Offer {
	var offers []duffel.Offer
	var err error

	for retry := 0; retry <= maxRetries; retry++ {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Panic occurred on attempt %d: %v\n", retry+1, r)
				if retry < maxRetries {
					// Retry the request after a timeout
					time.Sleep(3 * time.Second)
				}
			}
		}()

		// Attempt to create the offer request
		offers, err = createOfferRequest(dfl, origin, destination, departureDate)
		if err == nil {
			// Request succeeded, break out of the retry loop
			break
		}

		fmt.Printf("Error on attempt %d: %v\n", retry+1, err)

		if retry < maxRetries {
			// Retry the request after a timeout (you can adjust the timeout duration)
			time.Sleep(3 * time.Second)
		}
	}

	if err != nil {
		// Handle the final error here or return an empty offers slice
		fmt.Printf("All retry attempts failed: %v\n", err)
	}

	return offers
}

func createOfferRequest(dfl duffelClient, origin string, destination string, departureDate time.Time) ([]duffel.Offer, error) {
	maxConn := 0

	input := duffel.OfferRequestInput{
		Passengers: []duffel.OfferRequestPassenger{{Age: 20}},
		Slices: []duffel.OfferRequestSlice{
			{Origin: origin, Destination: destination, DepartureDate: duffel.Date(departureDate)},
		},
		CabinClass:      "economy",
		MaxConnections:  &maxConn,
		ReturnOffers:    true,
		SupplierTimeout: 10000,
	}

	request, err := dfl.CreateOfferRequest(context.Background(), input)
	if err != nil {
		panic(err)
	}

	return request.Offers, nil
}
