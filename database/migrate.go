package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Read PostgreSQL credentials from environment variables
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDBName := os.Getenv("POSTGRES_DB")
	pgSSLMode := os.Getenv("POSTGRES_SSLMODE")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", pgUser, pgPassword, pgDBName, pgSSLMode)

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Read the SQL script from the schema.sql file
	sqlFile, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Execute the SQL script
	_, err = db.Exec(string(sqlFile))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Schema.sql executed successfully.")
}
