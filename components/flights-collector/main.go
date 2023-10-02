package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/cors"
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
	conn, ch, err := getRabbitClient()
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

	declareQueue(ch, OfferSearchQueue)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow requests from any origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(r)

	r.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
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
		log.Fatal(http.ListenAndServe(":8080", handler))
	}()

	<-stopChan
}
