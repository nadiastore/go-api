package cache

import (
	"os"
	"strconv"

	// TODO: Buat reponya biar bisa diinstall
	"github.com/nadiastore/go-api/pkg/utils"

	"github.com/go-redis/redis/v8"
)

func RedisConnection() (*redis.Client, error) {
	dbNumber, _ := strconv.Atoi(os.Getenv("REDIS_DB_NUMBER"))
	redisConnURL, err := utils.ConnectionURLBuilder("redis")

	if err != nil {
		return nil, err
	}

	options := &redis.Options{
		Addr:     redisConnURL,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNumber,
	}

	return redis.NewClient(options), nil
}
