package main

import (
	"context"
	"github.com/airheartdev/duffel"
	"time"
)

type duffelClient interface {
	CreateOfferRequest(ctx context.Context, input duffel.OfferRequestInput) (*duffel.OfferRequest, error)
}

func getOffers(dfl duffelClient, origin string, destination string, departureDate time.Time) []duffel.Offer {
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

	return request.Offers
}
