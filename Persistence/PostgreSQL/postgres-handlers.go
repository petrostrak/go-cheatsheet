package main

import (
	"fmt"
	"log"
	"net/http"
	db "postgres/db/sqlc"
)

func (s *Server) createPerson(w http.ResponseWriter, r *http.Request) {
	log.Println("createPerson() invoked!")

	arg := db.CreatePersonParams{
		Kind:                   "Human",
		PersonsName:            "Petros Trak",
		Origins:                "Athens, Greece",
		ProgrammingLanguages:   []string{"Golang", "Java", "Javascript", "Rust"},
		Tools:                  []string{"Debian Linux", "Docker", "!# Bash", "MySQL", "Postgresql", "Redis"},
		Github:                 "https://github.com/petrostrak",
		Linkedin:               "https://www.linkedin.com/in/petrostrak/",
		Personal:               "https://petrostrak.netlify.app/",
		ForeignLanguages:       []string{"Greek", "English", "German"},
		FavFood:                "Ramen",
		FavDrink:               "Gin",
		FavProgrammingLanguage: "Golang",
		ThinkingAbout:          []string{"gRPC", "Concurrency in Go", "русский язык"},
		Hobbies:                []string{"Coding", "Foreign Languages", "Video Games"},
	}

	person, err := s.store.CreatePerson(r.Context(), arg)
	if err != nil {
		Error500(w, r)
		return
	}

	_ = WriteJson(w, http.StatusCreated, person)
}

func (s *Server) readAllPersons(w http.ResponseWriter, r *http.Request) {
	log.Println("readAllPersons() invoked!")

	persons, err := s.store.ListPersons(r.Context(), 5)
	if err != nil {
		Error500(w, r)
		return
	}

	_ = WriteJson(w, http.StatusCreated, persons)
}

func (s *Server) updatePerson(w http.ResponseWriter, r *http.Request) {}

func (s *Server) deletePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("deletePerson() invoked!")

	var id int64 = 0
	err := s.store.DeletePersonById(r.Context(), id)
	if err != nil {
		Error500(w, r)
		return
	}

	_ = WriteJson(w, http.StatusCreated, fmt.Sprintf("Deleted person with id of %d!", id))
}
