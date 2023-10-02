module github.com/prodigeris/Flight-searcher-go/database

go 1.21

require (
	github.com/lib/pq v1.10.9
	github.com/prodigeris/Flight-searcher-go/common v0.0.0
)

require github.com/rabbitmq/amqp091-go v1.8.1 // indirect

replace github.com/prodigeris/Flight-searcher-go/common v0.0.0 => ../common
