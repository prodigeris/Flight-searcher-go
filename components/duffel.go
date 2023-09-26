package components

import (
	"context"
	"github.com/airheartdev/duffel"
	"time"
)

type duffelClient interface {
	CreateOfferRequest(ctx context.Context, input duffel.OfferRequestInput) (*duffel.OfferRequest, error)
}

func getOffers(dfl duffelClient, origin string, destination string, departureDate time.Time, returnDate time.Time) []duffel.Offer {
	//dfl := duffel.New(os.Getenv("DUFFEL_TOKEN"))
	maxConn := 0

	input := duffel.OfferRequestInput{
		Passengers: []duffel.OfferRequestPassenger{{Type: "adult"}},
		Slices: []duffel.OfferRequestSlice{
			{Origin: origin, Destination: destination, DepartureDate: duffel.Date(departureDate)},
			{Origin: destination, Destination: origin, DepartureDate: duffel.Date(returnDate)},
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
