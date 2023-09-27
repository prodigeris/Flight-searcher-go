package main

import (
	"encoding/json"
	"fmt"
	"github.com/prodigeris/Flight-searcher-go/common"
	"log"
	"net/http"
)

func main() {

	db, err := common.GetDB()
	if err != nil {
		log.Fatalf("Failed to open connection to DB: %v", err)
	}

	http.HandleFunc("/itineraries", func(w http.ResponseWriter, r *http.Request) {

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

		w.Write(jsonResponse)
	})

	fmt.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
