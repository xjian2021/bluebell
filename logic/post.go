package logic

import (
	"fmt"
	"github.com/xjian2021/bluebell/dao/mysql"
	"github.com/xjian2021/bluebell/dao/redis"
	"github.com/xjian2021/bluebell/models"
	"github.com/xjian2021/bluebell/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(input *models.CreatePostInput) (err error) {
	postID := snowflake.GenID()
	newPost := &models.Post{
		ID:          postID,
		AuthorID:    input.AuthorID,
		CommunityID: input.CommunityID,
		Title:       input.Title,
		Content:     input.Content,
	}
	id, err := mysql.CreatePost(newPost)
	if err != nil {
		return fmt.Errorf("mysql CreatePost newPost:%+v err:%s", newPost, err.Error())
	}

	err = redis.CreatePost(postID)
	if err != nil {
		return fmt.Errorf("redis CreatePost newPost:%+v err:%s", newPost, err.Error())
	}
	zap.S().Infof("new post id:%d", id)
	return nil
}

func PostDetail(id int64) (output *models.PostDetailResData, err error) {
	return mysql.GetPostDetail(id)
}

func PostList(input *models.PostListInput) (output []*models.Post, err error) {
	postIDs, err := redis.GetPostIdListByOrder(input.Order, input.LastPostID, input.Limit)
	if err != nil {
		return nil, fmt.Errorf("redis GetPostIdListByOrder err:%s", err.Error())
	}
	//postIDs := make([]int64, len(postIDsStr), len(postIDsStr))
	//for i, postIDStr := range postIDsStr {
	//	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	//	if err != nil {
	//		return nil, fmt.Errorf("ParseInt err:%s", err.Error())
	//	}
	//	postIDs[i] = postID
	//}
	return mysql.PostList(postIDs, input.Limit)
}
