package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/xjian2021/bluebell/middlewares"
	"github.com/xjian2021/bluebell/models"
	"github.com/xjian2021/bluebell/pkg/errorcode"
)

func Response(c *gin.Context, code errorcode.Code, msg string, data interface{}) {
	if msg == "" {
		msg = code.Error()
	}
	c.JSON(http.StatusOK, models.ResponseData{
		Code: code.Code(),
		Msg:  msg,
		Data: data,
	})
	c.Abort()
}

func HandleError(c *gin.Context, err error) {
	var (
		code = errorcode.CodeSuccess
		msg  string
	)
	if err != nil {
		zap.S().Errorf("\t%s -> handler fail err:%s", c.Value(middlewares.ReqKey), err.Error())
		if e, ok := err.(errorcode.Code); ok {
			code = e
		} else {
			code = errorcode.CodeUnknownError
			msg = "操作失败"
		}
	}
	Response(c, code, msg, nil)
}

func HandleOutput(c *gin.Context, output interface{}, err error) {
	var (
		code = errorcode.CodeSuccess
		msg  string
	)
	if err != nil {
		zap.S().Errorf("%s -> handler fail err:%s", c.Value(middlewares.ReqKey), err.Error())
		if e, ok := err.(errorcode.Code); ok {
			code = e
		} else {
			code = errorcode.CodeUnknownError
			msg = "操作失败"
		}
	}
	Response(c, code, msg, output)
}
