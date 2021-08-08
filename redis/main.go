package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var rdb *redis.Client

func initRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err = rdb.Ping(ctx).Result()
	return
}

func main() {
	if err := initRedis(); err != nil {
		fmt.Printf("connect redis failed err:%v\n", err)
		panic(err)
	}
	fmt.Println("connect redis success")
	if err := rdb.Set(ctx, "name", "value1", 0).Err(); err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key does not exist")
	case err != nil:
		fmt.Println("Get failed", err)
	case val2 == "":
		fmt.Println("value is empty")
	}
}
