package middlewares

import (
	"fmt"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

const (
	ReqKey = "req-key"
)

var apiID uint64

//LoadApiMeta 接口id融合日志方便调试
func LoadApiMeta(c *gin.Context) {
	c.Set(ReqKey, fmt.Sprintf("%s::%d", c.FullPath(), atomic.AddUint64(&apiID, 1)))
}

func JwtAuth(c *gin.Context) {
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的Authorization中，并使用Bearer开头
	// 这里的具体实现方式要依据你的实际业务情况决定
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {

	}
}
