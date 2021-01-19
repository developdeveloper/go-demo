package session

import "github.com/go-redis/redis"

func GetRedisClient() *redis.Client {
	return GetRedisClientWithDB(0)
}

func GetRedisClientWithDB(db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       db,
	})
}
