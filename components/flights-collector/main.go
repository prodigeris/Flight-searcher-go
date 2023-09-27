package main

import (
	"fmt"
	"github.com/prodigeris/Flight-searcher-go/common"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const InquiriesQueue = "inquiries"

func main() {
	conn, ch, err := common.GetRabbitClient()
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ client: %v", err)
	}
	defer conn.Close()
	defer ch.Close()

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
		fmt.Printf("Received a message: %s\n", msg.Body)

		// Simulate some processing time
		time.Sleep(2 * time.Second)

		// Acknowledge the message (manual ack)
		msg.Ack(false)
	}
}
