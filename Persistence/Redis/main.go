package main

import (
	"net/http"
)

var (
	SERVER_ADDRESS = "0.0.0.0:8001"
	REDIS_HOST     = "localhost:6379"
	REDIS_PASSWORD = ""
	REDIS_PREFIX   = "cheatsheet"
)

type Server struct {
	cache RedisCache
}

func (s *Server) setupRouter() {
	mux := http.NewServeMux()

	mux.HandleFunc("/save-in-cache", s.saveInCache)
	mux.HandleFunc("/get-from-cache", s.getFromCache)
	mux.HandleFunc("/delete-from-cache", s.deleteFromCache)
	mux.HandleFunc("/empty-cache", s.emptyCache)

	err := http.ListenAndServe(SERVER_ADDRESS, mux)
	if err != nil {
		panic(err)
	}
}

func main() {

	server := &Server{
		cache: RedisCache{
			Conn:   CreateClientRedisCache().Conn,
			Prefix: REDIS_PREFIX,
		},
	}

	server.setupRouter()
}
