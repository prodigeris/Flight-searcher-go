package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/prodigeris/Flight-searcher-go/common"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type Search struct {
	FromAirport string    `json:"from-airport"`
	ToAirport   string    `json:"to-airport"`
	Date        time.Time `json:"date"`
}

func launchSearches(ch *amqp.Channel, weekendCount int) {
	fmt.Println("Starting to consume inquiry for weekend count: ", weekendCount)
	dFlights, aFlights := GetAllDeparturesAndArrivals()
	friday, sunday := nextFridayAndSunday(time.Now())
	searches := make([]Search, 0)
	for _, flight := range dFlights {
		searches = append(searches, Search{FromAirport: flight.FromAirport, ToAirport: flight.ToAirport, Date: friday})
	}
	for _, flight := range aFlights {
		searches = append(searches, Search{FromAirport: flight.FromAirport, ToAirport: flight.ToAirport, Date: sunday})
	}
	for _, search := range searches {
		publishSearch(ch, search)
	}
}

func publishSearch(ch *amqp.Channel, search Search) {
	searchJson, err := json.Marshal(search)
	if err != nil {
		fmt.Printf("Error encoding Search to JSON: %v\n", err)
		return
	}
	err = ch.PublishWithContext(context.Background(), "", common.OfferSearchQueue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        searchJson,
	})
	if err != nil {
		fmt.Printf("Error publishing Search: %v\n", err)
	}
	fmt.Printf("Published search Search: %v\n", search)
}
