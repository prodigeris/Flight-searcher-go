# Flight Searcher - Go

## Overview
Flight Searcher is a robust application developed in Go, designed to aid users in discovering the cheapest flights and generating itineraries from Vilnius and Kaunas to all possible destinations, allowing users to browse and get inspired for their next trips.

## Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/prodigeris/Flight-searcher-go
   ```
2. Install the necessary dependencies and create a .env file from .env.example:

    ```
    make install
    ```
3. After running the install command, configure the .env file by adding your Duffel credentials.

    ```sh
    DUFFEL_TOKEN=<Your Duffel Token>
    ```
## Setup & Run
This project is configured to run multiple services using Docker Compose, including PostgreSQL, RabbitMQ, Prometheus, Grafana, and several Go components.

### Start Services:
Navigate to the project's root directory and start all services.

```
docker-compose up -d
```
This command will start all the services defined in the docker-compose.yml file in detached mode.


### Initialize Database:
The migrate service will automatically handle the database migrations.

### Accessing the Services:

Flight Searcher Application: http://127.0.0.1:8383/

Grafana: http://127.0.0.1:3331/

Prometheus: http://127.0.0.1:9090/

### Usage
Once the application is running, users can select how many weekends they want to generate the flights for and the app will create itineraries from Vilnius and Kaunas to all possible destinations. It fetches the cheapest flights for these destinations, allowing users to browse and be inspired about where to go next.

## Testing
Run the test cases to ensure the functionality of the application.
```
make test
```