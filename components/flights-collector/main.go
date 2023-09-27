package main

import (
	"encoding/json"
	"github.com/prodigeris/Flight-searcher-go/common"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Inquiry struct {
	WeekendCount int `json:"weekend_count"`
}

func main() {
	conn, ch, err := common.GetRabbitClient()
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ client: %v", err)
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {

		}
	}(ch)

	common.DeclareQueue(ch, common.InquiriesQueue)
	common.DeclareQueue(ch, common.OfferSearchQueue)

	msgs, err := ch.Consume(
		common.InquiriesQueue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	if err != nil {
		log.Fatalf("Failed to open connection to DB: %v", err)
	}

	for msg := range msgs {
		var inquiry Inquiry
		err := json.Unmarshal(msg.Body, &inquiry)
		if err != nil {
			log.Printf("Failed to unmarshal message body: %v", err)
		}
		launchSearches(ch, inquiry.WeekendCount)
	}
}
