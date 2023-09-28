package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prodigeris/Flight-searcher-go/common"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow requests from any origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(r)

	r.HandleFunc("/itineraries", func(w http.ResponseWriter, r *http.Request) {
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
	})

	fmt.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
