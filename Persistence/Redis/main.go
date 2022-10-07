package main

import (
	"net/http"

	"github.com/gomodule/redigo/redis"
)

var (
	SERVER_ADDRESS = "0.0.0.0:8001"
	REDIS_HOST     = "localhost:6379"
	REDIS_PASSWORD = ""
	REDIS_PREFIX   = "cheatsheet"
	myRedisCache   *RedisCache
	redisPool      *redis.Pool
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
	myRedisCache = CreateClientRedisCache()
	redisPool = myRedisCache.Conn

	server := &Server{
		cache: RedisCache{
			Conn:   redisPool,
			Prefix: REDIS_PREFIX,
		},
	}

	server.setupRouter()
}
