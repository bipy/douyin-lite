package queries

import (
	"douyin-lite/platform"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

type DouyinQuery struct {
	*sqlx.DB
}

type RedisQuery struct {
	*redis.Client
}

var (
	DouyinDB *DouyinQuery
	RedisDB  *RedisQuery
)

func init() {
	DouyinDB = &DouyinQuery{DB: platform.GetNewMySQLConn()}
	//RedisDB = &RedisQuery{Client: platform.GetNewRedisConn()}
}
