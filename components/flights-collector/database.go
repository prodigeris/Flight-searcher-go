package main

//func insertFlight(db *sql.DB, fromAirport, toAirport string, flightDate time.Time) error {
//	// Prepare the SQL statement
//	stmt, err := db.Prepare("INSERT INTO flights (from_airport, to_airport, flight_date, created_at) VALUES ($1, $2, $3, $4)")
//	if err != nil {
//		return err
//	}
//	defer stmt.Close()
//
//	// Execute the SQL statement to insert a new flight record
//	_, err = stmt.Exec(fromAirport, toAirport, flightDate, time.Now())
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
