package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/xjian2021/bluebell/logic"
	"github.com/xjian2021/bluebell/models"
	"github.com/xjian2021/bluebell/pkg/errorcode"
)

func SignUpHandler(c *gin.Context) {
	repID := c.Value(ReqKey)
	input := &models.SignUpInput{}
	if err := AuthBindJson(c, input); err != nil {
		zap.S().Errorf("%s -> BindJSON fail err:%s", repID, err.Error())
		Response(c, errorcode.CodeInvalidParam, err.Error(), nil)
		return
	}
	zap.S().Debugf("%s -> SignUpInput:%+v", repID, input)

	err := logic.SignUp(input)
	HandleError(c, err)
}

func LoginHandler(c *gin.Context) {
	repID := c.Value(ReqKey)
	input := &models.LoginInput{}
	if err := AuthBindJson(c, input); err != nil {
		zap.S().Errorf("%s -> BindJSON fail err:%s", repID, err.Error())
		Response(c, errorcode.CodeInvalidParam, err.Error(), nil)
		return
	}
	zap.S().Debugf("%s -> LoginInput:%+v", repID, input)

	output, err := logic.Login(input)
	HandleOutput(c, output, err)
}
