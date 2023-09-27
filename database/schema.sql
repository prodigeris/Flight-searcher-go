CREATE TABLE IF NOT EXISTS flights (
                                       id SERIAL PRIMARY KEY,
                                       from_airport VARCHAR(3) NOT NULL,
                                       to_airport VARCHAR(3) NOT NULL,
                                       flight_date DATE NOT NULL,
                                       updated_at TIMESTAMP NULL,
                                       created_at TIMESTAMP NOT NULL,
                                       UNIQUE (from_airport, to_airport, flight_date)
);

CREATE TABLE IF NOT EXISTS itineraries (
                                         id SERIAL PRIMARY KEY,
                                         departure_id INT,
                                         return_id INT,
                                         updated_at TIMESTAMP NULL,
                                         created_at TIMESTAMP NOT NULL,
                                         FOREIGN KEY (departure_id) REFERENCES flights(id),
                                         FOREIGN KEY (return_id) REFERENCES flights(id)
);
