## Go persistence with Redis

In this example, we persist data in Redis.

To connect to Redis:
```go
type redisConfig struct {
	host     string
	password string
	prefix   string
}

func createClientRedisCache() *cache.RedisCache {
	cacheClient := cache.RedisCache{
		Conn:   createRedisPool(),
		Prefix: redisConfig.prefix,
	}
	return &cacheClient
}

func createRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,
		MaxActive:   10000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",
				s.config.redis.host,
				redis.DialPassword(s.config.redis.password))
		},

		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}

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
```
Save in redis cache:
```go
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
```
Get from redis cache:
```go
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
```
Delete from redis cache:
```go
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
```
Empty redis cache:
```go
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
```