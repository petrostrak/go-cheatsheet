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

var (myRedisCache  *cache.RedisCache)

func main() {
    ...
    myRedisCache = createClientRedisCache()
}
```
Save in redis cache:
```go
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
```
Get from redis cache:
```go
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
```
Delete from redis cache:
```go
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
```
Empty redis cache:
```go
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
```