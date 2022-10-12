package middlewares

import (
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/xjian2021/bluebell/controller"
	"github.com/xjian2021/bluebell/pkg/errorcode"
	"github.com/xjian2021/bluebell/pkg/jwt"
)

var apiID uint64

//LoadApiMeta 接口id融合日志方便调试
func LoadApiMeta(c *gin.Context) {
	c.Set(controller.ReqKey, fmt.Sprintf("%s::%d", c.FullPath(), atomic.AddUint64(&apiID, 1)))
}

func JwtAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		zap.S().Debugf("%s -> Authorization", c.Value(controller.ReqKey))
		if authHeader == "" {
			controller.HandleError(c, errorcode.CodeInvalidToken)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			zap.S().Errorf("%s -> len(parts):%d parts[0]:%s", c.Value(controller.ReqKey), len(parts), parts[0])
			controller.HandleError(c, errorcode.CodeInvalidToken)
			return
		}

		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			zap.S().Errorf("%s -> ParseToken err:%s", c.Value(controller.ReqKey), err.Error())
			controller.HandleError(c, errorcode.CodeInvalidToken)
			return
		}

		c.Set(controller.UserIDKey, mc.UserID)
	}
}
