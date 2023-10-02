package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
)

type Search struct {
	FromAirport string    `json:"from-airport"`
	ToAirport   string    `json:"to-airport"`
	Date        time.Time `json:"date"`
}

func main() {
	conn, ch, err := getRabbitClient()
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

	declareQueue(ch, OfferSearchQueue)

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

	db, err := getDB()
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

func getDB() (*sql.DB, error) {
	pgHost := os.Getenv("PGHOST")
	pgPort := os.Getenv("PGPORT")
	pgUser := os.Getenv("PGUSER")
	pgPassword := os.Getenv("PGPASSWORD")
	pgDBName := os.Getenv("PGDATABASE")
	pgSSLMode := os.Getenv("PGDATABASE_SSLMODE")

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		pgHost, pgPort, pgUser, pgPassword, pgDBName, pgSSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return db, nil
}
