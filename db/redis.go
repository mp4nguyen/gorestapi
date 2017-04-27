package db

import (
	"fmt"

	"github.com/go-redis/redis"
)

var redisDB *redis.Client

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	redisDB = client
}

func GetRedis() *redis.Client {
	return redisDB
}
