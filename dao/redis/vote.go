package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/xjian2021/bluebell/pkg/errorcode"
	"math"
	"strconv"
	"time"
)

func CreatePost(postID int64) (err error) {
	pipeline := rdb.Pipeline()
	ctx := context.Background()
	pipeline.ZAdd(ctx, getPrefix(KeyPostTimeZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	pipeline.ZAdd(ctx, getPrefix(KeyPostScoreZSet), &redis.Z{
		Score:  0,
		Member: postID,
	})
	_, err = pipeline.Exec(ctx)
	return
}

func VoteForPost(userID, postID string, value float64) (err error) {
	var (
		ctx   = context.Background()
		vaKey = getPrefix(fmt.Sprintf(KeyVoteAction, postID))
		ptKey = getPrefix(KeyPostTimeZSet)
		psKey = getPrefix(KeyPostScoreZSet)
		ov    float64
		op    float64
	)
	//检查该文章是否还可以被投票
	validTime := rdb.ZScore(ctx, ptKey, postID).Val()
	now := time.Now()
	if float64(now.Unix())-validTime > validityTimeOfVoting.Seconds() {
		return errorcode.CodeVoteTimeExpired
	}
	//查出上一次对该文章的投票操作
	preAction := rdb.HGet(ctx, vaKey, userID).Val()
	if preAction != "" {
		ov, err = strconv.ParseFloat(preAction, 64)
		if err != nil {
			return err
		}
	}

	//根据上一次操作，来给本次操作进行结算(投票有效时间以及文章分数)
	var (
		diff     = math.Abs(ov - value) // 这里的好处是，哪怕用户不断重复操作，也不会有数据更新
		pipeline = rdb.Pipeline()
	)
	if ov > value {
		op = -1
	} else {
		op = 1
	}
	if diff*op == 0 {
		return nil
	}
	pipeline.ZIncrBy(ctx, ptKey, diff*op*scorePerVote, postID)
	// TODO zset的score不会去到负数？
	pipeline.ZIncrBy(ctx, psKey, diff*op, postID)
	//更新最后一次操作
	if value == 0 {
		pipeline.HDel(ctx, vaKey, userID)
	} else {
		pipeline.HSet(ctx, vaKey, userID, value)
	}
	_, err = pipeline.Exec(ctx)
	return
}
