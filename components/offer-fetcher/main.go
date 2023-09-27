package main

import (
	"encoding/json"
	"github.com/prodigeris/Flight-searcher-go/common"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const OfferSearchQueue = "offer-searches"

type Search struct {
	FromAirport string    `json:"from-airport"`
	ToAirport   string    `json:"to-airport"`
	Date        time.Time `json:"date"`
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

	common.DeclareQueue(ch, OfferSearchQueue)

	msgs, err := ch.Consume(
		OfferSearchQueue,
		"",
		false,
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

	db, err := common.GetDB()
	if err != nil {
		log.Fatalf("Failed to open connection to DB: %v", err)
	}

	for msg := range msgs {
		var search Search
		err := json.Unmarshal(msg.Body, &search)
		if err != nil {
			log.Printf("Failed to unmarshal message body: %v", err)
		}
		err = consumeSearch(db, search)
		if err != nil {
			log.Printf("Nacking a message with error: %v", err)
			err := msg.Nack(false, true)
			if err != nil {
				return
			}
		} else {
			log.Printf("Acking a message: %v", search)
			err = msg.Ack(false)
			if err != nil {
				return
			}
		}

	}
}
