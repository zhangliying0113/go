package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = redisdb.Ping(ctx).Result()
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Println("connect redis failed, err:", err)
	}
	fmt.Println("连接 redis 成功")

	key := "rank"
	items := []*redis.Z{
		&redis.Z{Score: 90, Member: "PHP"},
		&redis.Z{Score: 96, Member: "Golang"},
		&redis.Z{Score: 97, Member: "Python"},
		&redis.Z{Score: 99, Member: "Java"},
	}
	redisdb.ZAdd(context.Background(), key, items...).Result()

	newScorde, err := redisdb.ZIncrBy(context.Background(), key, 10, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScorde)
}
