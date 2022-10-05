package main

import (
	"net/http"
)

func saveInCache(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Name  string `json:"name"`
		Value string `json:"value"`
		CSRF  string `json:"csrf_token"`
	}

	err := ReadJSON(w, r, &userInput)
	if err != nil {
		Error500(w, r)
		return
	}

	err = Set(userInput.Name, userInput.Value)
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
}

func GetFromCache(w http.ResponseWriter, r *http.Request) {
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

	fromCache, err := Get(userInput.Name)
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

func DeleteFromCache(w http.ResponseWriter, r *http.Request) {
	var userInput struct {
		Name string `json:"name"`
		CSRF string `json:"csrf_token"`
	}

	err := ReadJSON(w, r, &userInput)
	if err != nil {
		Error500(w, r)
		return
	}

	err = Forget(userInput.Name)
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

func EmptyCache(w http.ResponseWriter, r *http.Request) {

	err := ReadJSON(w, r, &userInput)
	if err != nil {
		Error500(w, r)
		return
	}

	err = Empty()
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
