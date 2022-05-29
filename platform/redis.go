package platform

import (
	"douyin-lite/pkg/configs"
	"github.com/go-redis/redis/v8"
)

func GetNewRedisConn() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     configs.RedisHost,
		Password: configs.RedisPassword,
		DB:       0,
	})
}
