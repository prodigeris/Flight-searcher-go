package main

import (
	"encoding/json"
	"fmt"
	"github.com/prodigeris/Flight-searcher-go/common"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
		db, err := common.GetDB()
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
