package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

func main() {

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "43.154.141.194:6379",
		Password: "zah56641189", // no password set
		DB:       0,             // use default DB
	})

	val, _ := rdb.Get(ctx, "username").Result()
	fmt.Println(val)

}
