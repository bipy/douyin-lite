package platform

import (
	"douyin-lite/pkg/configs"
	"github.com/go-redis/redis/v8"
	"strconv"
)

func GetNewRedisConn() *redis.Client {
	dbNumber, _ := strconv.Atoi(configs.RedisDBNum)
	return redis.NewClient(&redis.Options{
		Addr:     configs.RedisHost,
		Password: configs.RedisPassword,
		DB:       dbNumber,
	})
}
