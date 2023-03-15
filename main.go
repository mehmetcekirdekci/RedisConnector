package main

import (
	"context"
	"fmt"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"time"
)

const redisCacheKey = "redisCacheKey"

func main() {
	connection := openConnection

	redisCache := prepareCache(connection())

	var objectToGotFromCache *User

	err := redisCache.Get(context.Background(), redisCacheKey, &objectToGotFromCache)
	if err != nil {
		fmt.Printf("There was an error in get. Error: %s", err.Error())
	}

	if objectToGotFromCache != nil {
		fmt.Printf("Object was got from the cache. Object: %v", objectToGotFromCache)
	} else {
		objectToBeCached := User{
			Name:     "Mehmet",
			LastName: "Çekirdekci",
			Age:      30,
			Country:  "Turkey",
		}

		item := new(cache.Item)
		item.Ctx = context.Background()
		item.Key = redisCacheKey
		item.Value = objectToBeCached
		item.TTL = time.Minute * 1

		err = redisCache.Set(item)
		if err != nil {
			fmt.Printf("There was an error in set. Error: %s", err.Error())
		} else {
			fmt.Print("Object was cached.")
		}
	}
}

func openConnection() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return rdb
}

func prepareCache(rdb *redis.Client) *cache.Cache {
	myCache := cache.New(&cache.Options{
		Redis: rdb,
	})

	return myCache
}

type User struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Age      int    `json:"age"`
	Country  string `json:"country"`
}
