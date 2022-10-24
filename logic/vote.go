package logic

import (
	"github.com/xjian2021/bluebell/dao/redis"
	"github.com/xjian2021/bluebell/models"
	"strconv"
)

/*
   投票算法：
   	赞 +1 踩 -1
   	使用redis的有序集合类型，以score大小来排序
   	记录每个人对指定文章的操作
   	每增加200赞相当于可以多一天的可投票天数
   	以可投票过期时间作为分数
   	一个赞分数：432

   	分析：
   	上一次没投，这次投赞		+1  abs(0-1) = 1
   	上一次投踩，这次弃票		+1  abs(-1-0) = 1
   	上一次投踩，这次投赞		+2 	abs(-1-1) = 2
   	上一次投赞，这次弃票		-1	abs(1-0) = 1 ?
   	上一次没投，这次投踩		-1	abs(0+1) = 1 ?
   	上一次投赞，这次投踩		-2	abs(1+1) = 2 ?

   	TODO 在添加文章的接口上要加上生成文章可投票时间的缓存数据
*/
func Vote(input *models.VoteInput) (err error) {
	return redis.VoteForPost(strconv.Itoa(int(input.UserID)), input.PostID, input.Operating)
}
