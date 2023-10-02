package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting the migration")

	db, err := getDB()
	if err != nil {
		log.Fatal("Cannot connect to the database", err)
	}

	sqlFile, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration has been successful.")
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
