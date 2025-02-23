// redis.go
package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

// In redis.go - Add connection verification
func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"), // Use Redis port for Order Service
		Password: "",
		DB:       0,
	})

	// Add retry logic
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("‼️ Failed to connect to Redis: %v", err) // More visible error
	}

	return rdb
}
