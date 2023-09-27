package common

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func GetDB() (*sql.DB, error) {
	pgHost := os.Getenv("POSTGRES_HOST")
	pgPort := os.Getenv("POSTGRES_PORT")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDBName := os.Getenv("POSTGRES_DB")
	pgSSLMode := os.Getenv("POSTGRES_SSLMODE")

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
