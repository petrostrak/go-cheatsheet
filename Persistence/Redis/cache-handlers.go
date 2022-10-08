package main

import (
	"log"
	"net/http"
)

var (
	userInput = Person{
		Kind: "Human",
		Metadata: Metadata{
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
		Favorites: Favorites{
			Food:           "Ramen",
			Drink:          "Gin",
			ProgramingLang: "Golang",
		},
		ThinkingAbout: []string{
			"gRPC",
			"Concurrency in Go",
			"русский язык",
		},
		Hobbies: []string{
			"Coding",
			"Foreign Languages",
			"Video Games",
		},
	}
)

func (s *Server) saveInCache(w http.ResponseWriter, r *http.Request) {
	log.Println("saveInCache() invoked!")

	err := s.cache.Set(userInput.Metadata.Name, userInput.Metadata.Locations.Github)
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

	fromCache, err := s.cache.Get(userInput.Metadata.Name)
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
	log.Println("deleteFromCache() invoked!")

	err := s.cache.Forget(userInput.Metadata.Name)
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
	log.Println("emptyCache() invoked!")

	err := s.cache.Empty()
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
