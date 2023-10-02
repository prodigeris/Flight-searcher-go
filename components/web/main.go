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
	fmt.Println("Starting WEB service")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
