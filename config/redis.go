package config

import (
	"context"
	"github.com/apex/log"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
)

func InitRedis() (client *redis.Client) {
	ctx := context.Background()
	address := os.Getenv("REDIS_URL")
	password := os.Getenv("REDIS_PASSWORD")
	index, _ := strconv.Atoi(os.Getenv("REDIS_INDEX"))

	client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       index,
	})

	client.Options().OnConnect = func(ctx context.Context, cn *redis.Conn) error {
		return cn.ClientSetName(ctx, "api-process-pdf").Err()
	}

	status := client.Ping(ctx)
	if status.Err() != nil {
		log.Fatalf("[InitRedis] - Erro na conexão com Redis - %s", status.Err())
	}
	return
}
