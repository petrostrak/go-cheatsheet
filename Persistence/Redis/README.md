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