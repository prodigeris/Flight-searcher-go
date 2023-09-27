package main

import (
	"fmt"
	"os"
)

func getClient() {
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDBName := os.Getenv("POSTGRES_DB")
	pgSSLMode := os.Getenv("POSTGRES_SSLMODE")
	fmt.Println(pgUser, pgPassword, pgDBName, pgSSLMode)
}
