package main

import (
	"encoding/json"
	"github.com/prodigeris/Flight-searcher-go/common"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
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
			log.Printf("Failed to close RabbitMQ connection: %v", err)
		}
	}(conn)
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			log.Printf("Failed to close RabbitMQ channel: %v", err)
		}
	}(ch)

	common.DeclareQueue(ch, common.OfferSearchQueue)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		// Ensure the request method is POST
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method. Only POST is allowed.", http.StatusMethodNotAllowed)
			return
		}

		// Parse the request JSON
		var inquiry Inquiry
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&inquiry); err != nil {
			http.Error(w, "Failed to decode JSON request", http.StatusBadRequest)
			return
		}

		launchSearches(ch, inquiry.WeekendCount)
		w.WriteHeader(http.StatusAccepted)
	})

	go func() {
		// Start the HTTP server
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	<-stopChan
}
