package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/xjian2021/bluebell/models"
	"github.com/xjian2021/bluebell/pkg/errorcode"
)

const (
	ReqKey  = "req-key"
	CodeKey = "code-key"
	MsgKey  = "msg-key"
	DataKey = "data-keyt"
)

func Response(c *gin.Context) {
	data := c.Value(DataKey)
	code := c.Value(CodeKey).(errorcode.Code)
	msg, exits := c.Get(MsgKey)
	if !exits || msg == "" {
		msg = code.Error()
	}
	c.JSON(http.StatusOK, models.ResponseData{
		Code: code.Code(),
		Msg:  msg,
		Data: data,
	})
}

func HandleError(c *gin.Context, err error) {
	if err != nil {
		zap.S().Errorf("%s -> handler fail err:%s", c.Value(ReqKey), err.Error())
		if e, ok := err.(errorcode.Code); ok {
			c.Set(CodeKey, e)
		} else {
			SetCodeMsg(c, errorcode.CodeUnknownError, "操作失败")
		}
		return
	}
	c.Set(CodeKey, errorcode.CodeSuccess)
}

func HandleOutput(c *gin.Context, output interface{}, err error) {
	var (
		errCode = errorcode.CodeSuccess
		msg     string
	)
	if err != nil {
		zap.S().Errorf("%s -> handler fail err:%s", c.Value(ReqKey), err.Error())
		if e, ok := err.(errorcode.Code); ok {
			errCode = e
		} else {
			errCode = errorcode.CodeUnknownError
			msg = "操作失败"
		}
	}
	SetCodeMsgData(c, errCode, msg, output)
}

func SetCodeMsg(c *gin.Context, code errorcode.Code, msg interface{}) {
	c.Set(CodeKey, code)
	c.Set(MsgKey, msg)
}

func SetCodeData(c *gin.Context, code errorcode.Code, data interface{}) {
	c.Set(CodeKey, code)
	c.Set(DataKey, data)
}
func SetCodeMsgData(c *gin.Context, code errorcode.Code, msg, data interface{}) {
	c.Set(CodeKey, code)
	c.Set(MsgKey, msg)
	c.Set(DataKey, data)
}
