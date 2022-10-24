package logic

import (
	"fmt"
	"github.com/xjian2021/bluebell/dao/redis"
	"go.uber.org/zap"

	"github.com/xjian2021/bluebell/dao/mysql"
	"github.com/xjian2021/bluebell/models"
	"github.com/xjian2021/bluebell/pkg/snowflake"
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

	err = redis.CreatePost(id)
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
	return mysql.PostList(input.LastPostID, input.Limit)
}
