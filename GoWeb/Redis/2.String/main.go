package main

import (
	"context"
	"fmt"
	"time"

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

	// Set
	err := rdb.Set(ctx, "gorediskey", "goredisvalue", 0).Err()
	if err != nil {
		panic(err)
	}

	// Get
	value, err := rdb.Get(ctx, "gorediskey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("gorediskey", value)

	// GetSet
	oldVal, err := rdb.GetSet(ctx, "gorediskey", "new value").Result()
	if err != nil {
		panic(err)
	}

	// old keyvalue
	fmt.Println("key:", oldVal)

	// SetNX
	err = rdb.SetNX(ctx, "key1", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// MSet
	err = rdb.MSet(ctx, "key1", "value1", "key2", "value2", "key3", "value3").Err()
	if err != nil {
		panic(err)
	}

	// MGet
	vals, err := rdb.MGet(ctx, "key1", "key2", "key3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals)

	// Incr函数每次加一
	val, err := rdb.Incr(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Incr最新值:", val)

	// IncrBy函数，可以指定每次递增数量
	valBy, err := rdb.IncrBy(ctx, "key", 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("IncrBy最新值:", valBy)

	// IncrByFloat函数，以浮点数的方式指定每次递增数量
	valFloat, err := rdb.IncrByFloat(ctx, "key", 2.2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("IncrByFloat最新值:", valFloat)

	// Decr函数每次减一
	val, err = rdb.Decr(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Decr最新值:", val)

	// DecrBy函数，可以指定每次递减数量
	valBy, err = rdb.DecrBy(ctx, "key", 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("DecrBy最新值:", valBy)

	// Del 删除，支持批量删除
	rdb.Del(ctx, "key")

	// 批量删除
	err = rdb.Del(ctx, "key1", "key2", "key3").Err()
	if err != nil {
		panic(err)
	}

	// Expire, 设置过期时间
	rdb.Expire(ctx, "key", 3*time.Second)

}
