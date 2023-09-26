package components

import (
	"context"
	"github.com/airheartdev/duffel"
	"testing"
	"time"
)

type mockDuffelClient struct{}

func (m *mockDuffelClient) CreateOfferRequest(ctx context.Context, input duffel.OfferRequestInput) (*duffel.OfferRequest, error) {
	return &duffel.OfferRequest{
		Offers: []duffel.Offer{
			{
				ID:               "123",
				LiveMode:         true,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
				ExpiresAt:        time.Now().Add(24 * time.Hour),
				TotalEmissionsKg: "10",
				RawTotalCurrency: "USD",
				RawTotalAmount:   "200",
				RawTaxAmount:     "20",
				RawTaxCurrency:   "USD",
				RawBaseAmount:    "180",
				RawBaseCurrency:  "USD",
			},
		},
	}, nil
}

func TestGetOffersReturnsExpectedOffers(t *testing.T) {
	mockClient := &mockDuffelClient{}

	origin := "JFK"
	destination := "LAX"
	departureDate := time.Now()
	returnDate := departureDate.Add(7 * 24 * time.Hour) // 7 days later

	offers := getOffers(mockClient, origin, destination, departureDate, returnDate)

	if len(offers) == 0 {
		t.Errorf("Expected offers to be non-empty, but got empty offers")
	}

	expectedID := "123"
	expectedTotalEmissionsKg := "10"

	for i, offer := range offers {
		if offer.ID != expectedID {
			t.Errorf("Offer %d: Expected ID %s, got %s", i, expectedID, offer.ID)
		}
		if offer.TotalEmissionsKg != expectedTotalEmissionsKg {
			t.Errorf("Offer %d: Expected TotalEmissionsKg %s, got %s", i, expectedTotalEmissionsKg, offer.TotalEmissionsKg)
		}
	}
}
