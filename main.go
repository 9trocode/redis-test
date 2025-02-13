package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Create a background context for Redis operations.
	ctx := context.Background()

	// Create a new Redis client.
	// Adjust the Addr and Password fields as needed.
	rdb := redis.NewClient(&redis.Options{
		Addr:     "lingering-shadow.big-action-beta.svc.pipeops.internal:6379",
		Password: "d7339e7366ab5e240e048ee57", // set your Redis password here
		DB:       0,                           // use default DB
	})

	// Ping the Redis server to test the connection.
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	fmt.Printf("Connected to Redis: %s\n", pong)

	time.Sleep(30 * time.Minute)
}
