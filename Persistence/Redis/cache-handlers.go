package main

import (
	"log"
	"net/http"
)

func (s *Server) saveInCache(w http.ResponseWriter, r *http.Request) {
	log.Println("saveInCache() invoked!")

	userInput := Person{
		kind: "Human",
		metadata: Metadata{
			name: "Petros Trakadas",
			from: "🇬🇷",
			programmingLanguages: []string{
				"Golang",
				"Java",
				"Javascript",
				"Rust",
			},
			tools: []string{
				"Debian Linux",
				"Docker",
				"!# Bash",
				"MySQL",
				"Postgresql",
				"Redis",
			},
			locations: Locations{
				github:   "https://github.com/petrostrak",
				linkedin: "https://www.linkedin.com/in/petrostrak/",
				personal: "https://petrostrak.netlify.app/",
			},
			foreignLanguages: []string{
				"🇬🇷",
				"🏴󠁧󠁢󠁥󠁮󠁧󠁿",
				"🇩🇪",
			},
		},
		favorites: Favorites{
			food:           "🍣",
			drink:          "🍺",
			programingLang: "Golang",
		},
		thinkingAbout: []string{
			"gRPC",
			"Concurrency in Go",
			"русский язык",
		},
		hobbies: []string{
			"Coding",
			"Foreign Languages",
			"🎮",
		},
	}

	log.Println(userInput)

	err := ReadJSON(w, r, &userInput)
	if err != nil {
		Error500(w, r)
		return
	}

	err = s.cache.Set(userInput.metadata.name, userInput.metadata.locations.github)
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
	var msg string
	var inCache = true

	var userInput struct {
		Name string `json:"name"`
		CSRF string `json:"csrf_token"`
	}

	err := ReadJSON(w, r, &userInput)
	if err != nil {
		Error500(w, r)
		return
	}

	fromCache, err := s.cache.Get(userInput.Name)
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
	var userInput struct {
		Name string `json:"name"`
		CSRF string `json:"csrf_token"`
	}

	err := ReadJSON(w, r, &userInput)
	if err != nil {
		Error500(w, r)
		return
	}

	err = s.cache.Forget(userInput.Name)
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
