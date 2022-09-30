package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"go.uber.org/zap"

	"github.com/xjian2021/bluebell/settings"
)

var rdb *redis.Client

func Init(c *settings.Redis) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       c.Db,
		PoolSize: c.PoolSize,
	})
	if _, err := rdb.Ping().Result(); err != nil {
		panic(fmt.Errorf("rdb.Ping.Result fail err:%s", err.Error()))
	}
	zap.L().Info("redis init success...")
}

func Close() {
	_ = rdb.Close()
}
