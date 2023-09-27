package main

import (
	"database/sql"
	"time"
)

type Flight struct {
	ID          int       `db:"id"`
	FromAirport string    `db:"from_airport"`
	ToAirport   string    `db:"to_airport"`
	Price       int       `db:"price"`
	FlightDate  time.Time `db:"flight_date"`
}

type Itinerary struct {
	DepartureFlight Flight
	ReturnFlight    Flight
	TotalPrice      int
}

func getFlights(db *sql.DB) ([]Flight, error) {
	query := "SELECT id, from_airport, to_airport, price, flight_date FROM flights"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var flights []Flight

	for rows.Next() {
		var flight Flight
		err := rows.Scan(
			&flight.ID,
			&flight.FromAirport,
			&flight.ToAirport,
			&flight.Price,
			&flight.FlightDate,
		)
		if err != nil {
			return nil, err
		}
		flights = append(flights, flight)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return flights, nil
}
