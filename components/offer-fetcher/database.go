package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func insertFlight(db *sql.DB, fromAirport, toAirport string, price int, flightDate time.Time) error {
	stmt, err := db.Prepare("INSERT INTO flights (from_airport, to_airport, price, flight_date, created_at) VALUES ($1, $2, $3, $4, $5) ON CONFLICT (from_airport, to_airport, flight_date) DO UPDATE SET price = $3, updated_at = $5")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Failed to close transaction")
		}
	}(stmt)

	_, err = stmt.Exec(fromAirport, toAirport, price, flightDate, time.Now())
	if err != nil {
		return err
	}

	return nil
}
