package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"

	"github.com/xjian2021/bluebell/logic"
	"github.com/xjian2021/bluebell/models"
	"github.com/xjian2021/bluebell/pkg/errorcode"
)

func CreatePostHandler(c *gin.Context) {
	repID := c.Value(ReqKey)
	input := &models.CreatePostInput{}
	if err := AuthBindJson(c, input); err != nil {
		zap.S().Errorf("%s -> BindJSON fail err:%s", repID, err.Error())
		Response(c, errorcode.CodeInvalidParam, err.Error(), nil)
		return
	}
	input.AuthorID = GetCurrenUserID(c)
	zap.S().Debugf("%s -> CreatePost:%+v", repID, input)

	err := logic.CreatePost(input)
	HandleError(c, err)
}

func PostDetailHandler(c *gin.Context) {
	repID := c.Value(ReqKey)
	id := c.Param("id")
	postID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		zap.S().Errorf("%s -> ParseInt fail err:%s", repID, err.Error())
		Response(c, errorcode.CodeInvalidParam, err.Error(), nil)
		return
	}
	zap.S().Debugf("%s -> PostDetail:%d", repID, postID)

	output, err := logic.PostDetail(postID)
	HandleOutput(c, output, err)
}

func PostListHandler(c *gin.Context) {
	repID := c.Value(ReqKey)
	input := &models.PostListInput{}
	if err := AuthBindQuery(c, input); err != nil {
		zap.S().Errorf("%s -> BindJSON fail err:%s", repID, err.Error())
		Response(c, errorcode.CodeInvalidParam, err.Error(), nil)
		return
	}
	zap.S().Debugf("%s -> PostList:%+v", repID, input)
	output, err := logic.PostList(input)
	HandleOutput(c, output, err)
}
