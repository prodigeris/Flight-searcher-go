package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {

	// Load connection string from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

}
