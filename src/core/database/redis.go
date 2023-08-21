package database

import (
	"os"
	"runtime"
	"strconv"

	"github.com/gofiber/storage/redis/v2"
)

var redisClient *RedisClient

type RedisClient struct {
	store *redis.Storage
}

// GetRedisClient for getting the only instance of the redis client
func GetRedisClient() *RedisClient {
	if redisClient == nil {
		redisClient = &RedisClient{}
		redisClient.Connect()
	}
	return redisClient
}

// Connect func for connect to redis
func (rc *RedisClient) Connect() {
	port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		port = 6379
	}

	// Initialize custom config
	store := redis.New(redis.Config{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     port,
		Username: os.Getenv("REDIS_USER"),
		Password: os.Getenv("REDIS_PASSWORD"),
		PoolSize: 10 * runtime.GOMAXPROCS(0),
	})

	// Set store globally
	rc.store = store
}

// Get func for get value from redis
func (rc *RedisClient) Get(key string) (string, error) {
	res, err := rc.store.Get(key)
	return string(res), err
}

// Set func for set value in redis
func (rc *RedisClient) Set(key string, value string) error {
	return rc.store.Set(key, []byte(value), 0)
}

// Delete func for delete value in redis
func (rc *RedisClient) Delete(key string) error {
	return rc.store.Delete(key)
}

// Reset func for delete all values in redis
func (rc *RedisClient) Reset() error {
	return rc.store.Reset()
}
