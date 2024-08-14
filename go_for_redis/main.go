package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pong, err := rdb.Ping(ctx).Result()
	fmt.Println("Ping:", pong, err)

	err = rdb.Set(ctx, "name", "Hasan", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "name").Result()
	fmt.Println("Name:", val)

	err = rdb.HSet(ctx, "user:1", map[string]interface{}{
		"name": "Ali Koç",
		"age":  21,
		"city": "İstanbul",
	}).Err()
	if err != nil {
		panic(err)
	}
	m, err := rdb.HGetAll(ctx, "user:1").Result()
	fmt.Println("User:", m)

	err = rdb.LPush(ctx, "my_list", "apple", "banana", "orange").Err()
	if err != nil {
		panic(err)
	}
	elements, err := rdb.LRange(ctx, "my_list", 0, -1).Result()
	fmt.Println("List:", elements)

	err = rdb.SAdd(ctx, "my_set", "apple", "banana", "orange").Err()
	if err != nil {
		panic(err)
	}
	members, err := rdb.SMembers(ctx, "my_set").Result()
	fmt.Println("Set:", members)

	err = rdb.ZAdd(ctx, "my_zset", &redis.Z{Score: 1.0, Member: "apple"}, &redis.Z{Score: 2.0, Member: "banana"}).Err()
	if err != nil {
		panic(err)
	}
	members, err = rdb.ZRange(ctx, "my_zset", 0, -1).Result()
	fmt.Println("Sorted Set:", members)

}
