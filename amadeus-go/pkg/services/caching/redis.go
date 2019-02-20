package caching

import (
	"github.com/go-redis/redis"
	"time"
)

func InitRedis() (redisClient *redis.Client, err error) {
	redisClient = redis.NewClient(
		&redis.Options{
			Addr:     "redis-server:6379",
			Password: "",
			DB:       0,
		},
	)

	_, err = redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}

	return
}

func ReadFromRedis(redisClient *redis.Client, key string) (value string, ttl time.Duration, err error) {
	value, err = redisClient.Get(key).Result()
	if err != nil {
		return
	}

	ttl, err = redisClient.TTL(key).Result()
	if err != nil {
		return
	}

	return
}

func WriteToRedis(redisClient *redis.Client, key string, value string, ttl time.Duration) error {
	_, err := redisClient.Set(key, value, time.Second*ttl).Result()
	if err != nil {
		return err
	}

	return nil
}
