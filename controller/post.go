package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/xjian2021/bluebell/logic"
	"github.com/xjian2021/bluebell/models"
	"github.com/xjian2021/bluebell/pkg/errorcode"
)

func CreatePostHandler(c *gin.Context) {
	repID := c.Value(ReqKey)
	input := &models.CreatePost{}
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

func PostListHandler(c *gin.Context) {

}
