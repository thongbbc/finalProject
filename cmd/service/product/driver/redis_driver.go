package driver

import (
	"fmt"
	"github.com/go-redis/redis"
)

type RedisDB struct {
	DB *redis.Client
}

var Redis = &RedisDB{}

func ConnectRedis(redisHost string, port string) *RedisDB {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, port),
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		panic(err)
	}
	Redis.DB = client
	return Redis
}
