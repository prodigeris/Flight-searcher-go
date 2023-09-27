package main

import (
	"fmt"
	"github.com/prodigeris/Flight-searcher-go/common"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Starting the migration")

	db, err := common.GetDB()
	if err != nil {
		log.Fatal("Cannot connect to the database")
	}

	sqlFile, err := os.ReadFile("database/schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration has been successful.")
}
