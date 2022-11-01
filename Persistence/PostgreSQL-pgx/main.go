package main

import (
	"flag"
	"log"
	"net/http"
	"postgresql-pgx/pkg/db"

	"github.com/go-chi/chi/v5"
)

type application struct {
	DSN string
	DB  db.PostgresConn
}

func main() {
	app := application{}

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5", "Postgres Connection")
	flag.Parse()

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = db.PostgresConn{DB: conn}

	mux := chi.NewRouter()

	// print out a msg
	log.Println("Starting server on port 8080...")

	// start the server
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
