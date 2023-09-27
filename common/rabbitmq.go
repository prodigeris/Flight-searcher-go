package common

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

const OfferSearchQueue = "offer-searches"
const InquiriesQueue = "inquiries"

func GetRabbitClient() (*amqp.Connection, *amqp.Channel, error) {
	rabbitMQURL := os.Getenv("RABBITMQ_URL")

	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return nil, nil, err
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return nil, nil, err
	}

	return conn, ch, nil
}

func DeclareQueue(ch *amqp.Channel, name string) {
	_, err := ch.QueueDeclare(
		name, // Queue name
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}
}
