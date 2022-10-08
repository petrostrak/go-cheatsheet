package main

import (
	"log"
	"net/http"
)

func (s *Server) createPerson(w http.ResponseWriter, r *http.Request) {
	log.Println("createPerson() invoked!")
}

func (s *Server) readPersons(w http.ResponseWriter, r *http.Request) {}

func (s *Server) updatePerson(w http.ResponseWriter, r *http.Request) {}

func (s *Server) deletePerson(w http.ResponseWriter, r *http.Request) {}
