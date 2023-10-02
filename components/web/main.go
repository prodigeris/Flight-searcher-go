package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Config struct {
	ItineraryAnalyzerHost string
	FlightsCollectorHost  string
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", index())

	fmt.Printf("Starting WEB service on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func index() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		config := Config{
			ItineraryAnalyzerHost: os.Getenv("ITINERARY_ANALYZER_HOST"),
			FlightsCollectorHost:  os.Getenv("FLIGHTS_COLLECTOR_HOST"),
		}
		tmpl, err := template.ParseFiles("components/web/static/index.html")
		if err != nil {
			http.Error(w, "Could not load template", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, config)
		if err != nil {
			panic(err)
		}
	}
}
