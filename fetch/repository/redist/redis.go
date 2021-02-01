package redist

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type redist struct {
	client *redis.Client
}

func NewRedisClient() *redist {
	address := os.Getenv("REDIS_SERVER")
	db := os.Getenv("REDIS_DB")
	dbint, _ := strconv.Atoi(db)

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		DB:       dbint,
		Password: "",
	})

	return &redist{client: client}
}

func (rds *redist) GetFloat(key string) (float64, error) {
	return rds.client.Get(key).Float64()
}

func (rds *redist) GetString(key string) string {
	return rds.client.Get(key).Val()
}

func (rds *redist) Set(key string, value interface{}, exp time.Duration) {
	err := rds.client.Set(key, value, exp).Err()
	if err != nil {
		log.Println(err)
	}
}
