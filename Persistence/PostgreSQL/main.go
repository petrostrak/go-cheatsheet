package main

import (
	"database/sql"
	"log"
	"net/http"
	db "postgres/db/sqlc"

	_ "github.com/lib/pq"
)

var (
	DB_DRIVER      = "postgres"
	DB_SOURCE      = "postgresql://postgres:secret@localhost:5432/cheatsheet?sslmode=disable"
	SERVER_ADDRESS = "0.0.0.0:8001"
)

type Server struct {
	store db.Querier
}

func (s *Server) setupRouter() {
	mux := http.NewServeMux()

	mux.HandleFunc("/create", s.createPerson)
	mux.HandleFunc("/read-all", s.readAllPersons)
	mux.HandleFunc("/update", s.updatePerson)
	mux.HandleFunc("/delete", s.deletePerson)

	err := http.ListenAndServe(SERVER_ADDRESS, mux)
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := sql.Open(DB_DRIVER, DB_SOURCE)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.New(conn)
	server := &Server{
		store: store,
	}

	server.setupRouter()
}
