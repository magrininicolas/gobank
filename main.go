package main

import (
	"log"

	"github.com/magrininicolas/gobank/api"
	"github.com/magrininicolas/gobank/db"
)

func main() {
	db, err := db.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewApiServer(":3000", db)
	server.Run()
}
