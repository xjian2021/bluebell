package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/xjian2021/bluebell/logic"
)

func CommunityHandler(c *gin.Context) {
	repID := c.Value(ReqKey)
	zap.S().Debugf("%s -> LoginInput", repID)
	output, err := logic.GetCommunityList()
	HandleOutput(c, output, err)
}
