package storage

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")
}

func SaveURL(code, original string) {
	err := rdb.Set(ctx, code, original, 24*time.Hour).Err()
	if err != nil {
		log.Fatalf("Could not save URL: %v", err)
	}
	log.Printf("Saved URL: %s -> %s", code, original)
}

func GetURL(code string) (string, bool) {
	val, err := rdb.Get(ctx, code).Result()
	if err == redis.Nil {
		log.Printf("URL not found for code: %s", code)
		return "", false
	} else if err != nil {
		log.Fatalf("Could not get URL: %v", err)
	}
	log.Printf("Retrieved URL: %s -> %s", code, val)
	return val, true
}
