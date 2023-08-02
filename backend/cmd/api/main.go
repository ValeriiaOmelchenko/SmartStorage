package main

import (
	"fmt"
	"log"
	"net/http"

	"smart_storage/internals/db"
)

const (
	port int = 8080
)

type application struct {
	DSN string
	DB  db.PostgresDB
}

func main() {
	app := application{
		DSN: "host=localhost port=5432 user=postgres password=password dbname=smart_storage sslmode=disable timezone=UTC connect_timeout=6",
	}
	log.Println("Starting application on port", port)

	// connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = db.PostgresDB{DB: conn}
	defer app.DB.Connection().Close()

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
