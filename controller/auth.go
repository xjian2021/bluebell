package controller

import (
	"fmt"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/xjian2021/bluebell/pkg/errorcode"
	validatorPkg "github.com/xjian2021/bluebell/pkg/validator-trans"
)

var apiID uint64

//LoadApiMeta 接口id融合日志方便调试
func LoadApiMeta(c *gin.Context) {
	c.Set(ReqKey, fmt.Sprintf("%s::%d", c.FullPath(), atomic.AddUint64(&apiID, 1)))
}

func AuthBindJson(c *gin.Context, input interface{}) (err error) {
	if err = c.ShouldBindJSON(input); err != nil {
		errors, ok := err.(validator.ValidationErrors)
		if ok {
			SetCodeMsg(c, errorcode.CodeInvalidParam, errors.Translate(validatorPkg.Trans))
			return
		}

		c.Set(CodeKey, errorcode.CodeInvalidParam)
		return
	}
	return nil
}
