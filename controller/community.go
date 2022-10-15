package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xjian2021/bluebell/pkg/errorcode"
	"go.uber.org/zap"
	"strconv"

	"github.com/xjian2021/bluebell/logic"
)

func CommunityHandler(c *gin.Context) {
	repID := c.Value(ReqKey)
	zap.S().Debugf("%s -> LoginInput", repID)
	output, err := logic.GetCommunityList()
	HandleOutput(c, output, err)
}

func CommunityDetailHandler(c *gin.Context) {
	repID := c.Value(ReqKey)
	idStr := c.Param("id")
	communityID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.S().Errorf("%s -> ParseInt fail err:%s", repID, err.Error())
		Response(c, errorcode.CodeInvalidParam, err.Error(), nil)
		return
	}
	zap.S().Debugf("%s -> CommunityDetai:%d", repID, communityID)
	output, err := logic.GetCommunityDetail(communityID)
	HandleOutput(c, output, err)
}
