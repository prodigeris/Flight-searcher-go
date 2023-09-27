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

const InquiriesQueue = "inquiries"

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

	common.DeclareQueue(ch, InquiriesQueue)

	msgs, err := ch.Consume(
		InquiriesQueue,
		"",
		false, // Auto Acknowledge (false for manual ack)
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// Handle incoming messages
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	for msg := range msgs {
		var inquiry Inquiry
		err := json.Unmarshal(msg.Body, &inquiry)
		if err != nil {
			log.Printf("Failed to unmarshal message body: %v", err)
		}
		go collect(inquiry.WeekendCount)
	}
}
