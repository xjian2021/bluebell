package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/xjian2021/bluebell/settings"
)

const (
	prefix = "test-"
	//KeyVoteAction 最后一次对文章投票的操作记录 redis的string类型
	KeyVoteAction = "vote:action::%s"
	//KeyPostTimeZSet 可投票时间 redis的ZSET类型 key:
	KeyPostTimeZSet = "post:time"
	//KeyPostScoreZSet 帖子与投票的分数
	KeyPostScoreZSet     = "post:score"
	validityTimeOfVoting = 24 * 7 * time.Hour
	scorePerVote         = 432
)

var rdb *redis.Client

func Init(c *settings.Redis) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       c.Db,
		PoolSize: c.PoolSize,
	})
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		panic(fmt.Errorf("rdb.Ping.Result fail err:%s", err.Error()))
	}
	zap.L().Info("redis init success...")
}

func getPrefix(key string) string {
	return prefix + key
}

func Close() {
	_ = rdb.Close()
}
