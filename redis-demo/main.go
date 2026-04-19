package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		PoolSize: 100,
	})

	_, err = rdb.Ping().Result()
	return err
}

func redisExample() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:=%v\n", err)
		return
	}

	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed,err:%v\n",err)
		return
	} else {
		fmt.Println("name", val2)
	}
}
func main() {
	if err := initClient(); err != nil {
		fmt.Printf("connect redis server failed,err:=%v\n", err)
		return
	}
	defer rdb.Close()
	fmt.Println("connect redis server success..")
	redisExample()
}
