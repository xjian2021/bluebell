package redis

import (
	"context"
	"strconv"

	"github.com/xjian2021/bluebell/models"
)

func GetPostIdListByOrder(order string, PrePostID, limit int64) ([]string, error) {
	ctx := context.Background()
	key := getPrefix(KeyPostTimeZSet)
	if order == models.OrderScore {
		key = getPrefix(KeyPostScoreZSet)
	}
	var start int64
	if PrePostID != 0 {
		start = rdb.ZRevRank(ctx, key, strconv.Itoa(int(PrePostID))).Val()
	}
	return rdb.ZRevRange(ctx, key, start, start+limit).Result()
}
