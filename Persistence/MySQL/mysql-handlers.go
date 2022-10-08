package main

import (
	"fmt"
	"log"
	db "mysql/db/sqlc"
	"net/http"
)

func (s *Server) createPerson(w http.ResponseWriter, r *http.Request) {
	log.Println("createPerson() invoked!")

	arg := db.CreatePersonParams{
		Kind:                   "Human",
		PersonsName:            "Petros Trak",
		Origins:                "Athens, Greece",
		ProgrammingLanguages:   "Golang, Java, Javascript ,Rust",
		Tools:                  "Debian Linux, Docker, !# Bash, MySQL, Postgresql, Redis",
		Github:                 "https://github.com/petrostrak",
		Linkedin:               "https://www.linkedin.com/in/petrostrak/",
		Personal:               "https://petrostrak.netlify.app/",
		ForeignLanguages:       "Greek, English, German",
		FavFood:                "Ramen",
		FavDrink:               "Gin",
		FavProgrammingLanguage: "Golang",
		ThinkingAbout:          "gRPC, Concurrency in Go, русский язык",
		Hobbies:                "Coding, Foreign Languages, Video Games",
	}

	_, err := s.store.CreatePerson(r.Context(), arg)
	if err != nil {
		log.Println(err)
		Error500(w, r)
		return
	}

	_ = WriteJson(w, http.StatusCreated, arg)
}

func (s *Server) readAllPersons(w http.ResponseWriter, r *http.Request) {
	log.Println("readAllPersons() invoked!")

	persons, err := s.store.ListPersons(r.Context())
	if err != nil {
		Error500(w, r)
		return
	}

	_ = WriteJson(w, http.StatusCreated, persons)
}

func (s *Server) updatePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("updatePerson() invoked!")

	arg := db.UpdatePersonParams{
		ID:                     1,
		Kind:                   "Alien",
		PersonsName:            "Petros Trak",
		Origins:                "Athens, Greece",
		ProgrammingLanguages:   "Golang, Java, Javascript ,Rust",
		Tools:                  "Debian Linux, Docker, !# Bash, MySQL, Postgresql, Redis",
		Github:                 "https://github.com/petrostrak",
		Linkedin:               "https://www.linkedin.com/in/petrostrak/",
		Personal:               "https://petrostrak.netlify.app/",
		ForeignLanguages:       "Greek, English, German",
		FavFood:                "Ramen",
		FavDrink:               "Gin",
		FavProgrammingLanguage: "Golang",
		ThinkingAbout:          "gRPC, Concurrency in Go, русский язык",
		Hobbies:                "Coding, Foreign Languages, Video Games",
	}

	_, err := s.store.UpdatePerson(r.Context(), arg)

	if err != nil {
		Error500(w, r)
		return
	}

	_ = WriteJson(w, http.StatusCreated, arg)
}

func (s *Server) deletePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("deletePerson() invoked!")

	var id int64 = 1
	err := s.store.DeletePersonById(r.Context(), id)
	if err != nil {
		Error500(w, r)
		return
	}

	_ = WriteJson(w, http.StatusCreated, fmt.Sprintf("Deleted person with id of %d!", id))
}
