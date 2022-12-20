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
	// user_1 是Hash key，username 是字段名，zhangsan 是字段值
	err := rdb.HSet(ctx, "user_1", "username", "zhangsan").Err()
	if err != nil {
		panic(err)
	}

	// user_1 是Hash key，username是字段名
	username, err := rdb.HGet(ctx, "user_1", "username").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(username)

	// 一次性返回key=user_1的所有hash字段和值
	data, err := rdb.HGetAll(ctx, "user_1").Result()
	if err != nil {
		panic(err)
	}

	// data是一个map类型，这里使用循环迭代输出
	for field, val := range data {
		fmt.Println(field, val)
	}

	// 累加count字段的值，一次性累加2，user_1为hash key
	count, err := rdb.HIncrBy(ctx, "user_1", "count", 2).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(count)

	// keys是一个string的数组
	keys, err := rdb.HKeys(ctx, "user_1").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(keys)

	// HLen
	size, err := rdb.HLen(ctx, "user_1").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(size)

	// HMGet支持多个field字段名，一次返回多个字段值
	vals, err := rdb.HMGet(ctx, "user_1", "username", "count").Result()
	if err != nil {
		panic(err)
	}

	// vals是一个数组
	fmt.Println(vals)

	// 初始化hash数据的多个字段值
	dataset := make(map[string]interface{})
	dataset["id"] = 1
	dataset["username"] = "lisi"

	// 一次性保存多个hash字段值
	err = rdb.HMSet(ctx, "key", dataset).Err()
	if err != nil {
		panic(err)
	}

	// HSetNX
	err = rdb.HSetNX(ctx, "key", "id", 100).Err()
	if err != nil {
		panic(err)
	}

	// HDel
	// 删除一个字段id
	rdb.HDel(ctx, "key", "id")

	// 删除多个字段
	rdb.HDel(ctx, "key", "id", "username")

	// 检测id字段是否存在
	err = rdb.HExists(ctx, "key", "id").Err()
	if err != nil {
		panic(err)
	}

}
