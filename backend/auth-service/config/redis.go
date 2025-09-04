package config

import "github.com/redis/go-redis/v9"

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // REDIS_HOST from env
		Password: "",               // Set if needed
		DB:       0,
	})
}
