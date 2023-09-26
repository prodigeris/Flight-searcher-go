package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDBName := os.Getenv("POSTGRES_DB")
	pgSSLMode := os.Getenv("POSTGRES_SSLMODE")

	// Define a request handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "hoiii!", pgUser, pgPassword, pgDBName, pgSSLMode)
		if err != nil {
			return
		}
	})

	// Specify the port to listen on
	port := 8080

	// Start the HTTP server
	fmt.Printf("Server is running on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
