package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()

	c := cors.New(corsOptions())

	handler := c.Handler(r)

	r.HandleFunc("/itineraries", itineraries())
	r.HandleFunc("/health", health())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server listening on : %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func health() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			return
		}
	}
}

func corsOptions() cors.Options {
	return cors.Options{
		AllowedOrigins: []string{"*"}, // Allow requests from any origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}
}

func itineraries() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := getDB()
		if err != nil {
			log.Fatalf("Failed to open connection to DB: %v", err)
		}

		flights, err := getFlights(db)
		if err != nil {
			log.Fatalf("Failed fetching flights: %v", err)
		}
		itineraries := generateItineraries(flights, 4)
		sortByPrice(itineraries)

		jsonResponse, err := json.Marshal(itineraries)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonResponse)
		if err != nil {
			return
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
