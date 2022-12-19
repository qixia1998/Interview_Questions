package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func main() {
	ctx := context.Background()
	//0代表永不过期
	err := rdb.Set(ctx, "gorediskey", "goredisvalue", 0).Err()
	if err != nil {
		panic(err)
	}
	value, err := rdb.Get(ctx, "gorediskey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("gorediskey", value)

	//OR
	val, err := rdb.Do(ctx, "get", "gorediskey").Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Println("gorediskey 不存在")
			return
		}
		panic(err)
	}
	fmt.Println("do operator : gorediskey", val.(string))
}
