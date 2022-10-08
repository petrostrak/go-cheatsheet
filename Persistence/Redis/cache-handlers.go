package main

import (
	"log"
	"net/http"
)

var (
	userInput = Person{
		kind: "Human",
		metadata: Metadata{
			Name: "Petros Trakadas",
			From: "Greece",
			ProgrammingLanguages: []string{
				"Golang",
				"Java",
				"Javascript",
				"Rust",
			},
			Tools: []string{
				"Debian Linux",
				"Docker",
				"!# Bash",
				"MySQL",
				"Postgresql",
				"Redis",
			},
			Locations: Locations{
				Github:   "https://github.com/petrostrak",
				Linkedin: "https://www.linkedin.com/in/petrostrak/",
				Personal: "https://petrostrak.netlify.app/",
			},
			ForeignLanguages: []string{
				"Greek",
				"English",
				"German",
			},
		},
		favorites: Favorites{
			Food:           "Ramen",
			Drink:          "Gin",
			ProgramingLang: "Golang",
		},
		thinkingAbout: []string{
			"gRPC",
			"Concurrency in Go",
			"русский язык",
		},
		hobbies: []string{
			"Coding",
			"Foreign Languages",
			"Video Games",
		},
	}
)

func (s *Server) saveInCache(w http.ResponseWriter, r *http.Request) {
	log.Println("saveInCache() invoked!")

	err := s.cache.Set(userInput.metadata.Name, userInput.metadata.Locations.Github)
	if err != nil {
		Error500(w, r)
		return
	}

	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	resp.Error = false
	resp.Message = "Saved in cache"

	_ = WriteJson(w, http.StatusCreated, resp)
	log.Println(resp)
}

func (s *Server) getFromCache(w http.ResponseWriter, r *http.Request) {
	log.Println("getFromCache() invoked!")
	var msg string
	var inCache = true

	fromCache, err := s.cache.Get(userInput.metadata.Name)
	if err != nil {
		msg = "Not found in cache!"
		inCache = false
	}

	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
		Value   string `json:"value"`
	}

	if inCache {
		resp.Error = false
		resp.Message = "success"
		resp.Value = fromCache.(string)
	} else {
		resp.Error = true
		resp.Message = msg
	}

	_ = WriteJson(w, http.StatusCreated, resp)
}

func (s *Server) deleteFromCache(w http.ResponseWriter, r *http.Request) {

	err := s.cache.Forget(userInput.metadata.Name)
	if err != nil {
		Error500(w, r)
		return
	}

	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	resp.Error = false
	resp.Message = "deleted from cache (if it existed)"

	_ = WriteJson(w, http.StatusCreated, resp)
}

func (s *Server) emptyCache(w http.ResponseWriter, r *http.Request) {
	var userInput struct{}

	err := ReadJSON(w, r, &userInput)
	if err != nil {
		Error500(w, r)
		return
	}

	err = s.cache.Empty()
	if err != nil {
		Error500(w, r)
		return
	}

	var resp struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	resp.Error = false
	resp.Message = "emptied cache"

	_ = WriteJson(w, http.StatusCreated, resp)
}
