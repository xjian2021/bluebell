package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xjian2021/bluebell/controller"
	"github.com/xjian2021/bluebell/logger"
	"github.com/xjian2021/bluebell/middlewares"
	"github.com/xjian2021/bluebell/pkg/snowflake"
	"github.com/xjian2021/bluebell/settings"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	g := r.Group("/api/v1")
	g.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "version:%s", settings.C.Version)
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{"msg_id": snowflake.GenID(), "msg": "api not found"},
		)
	})
	g.POST("/sign-up", middlewares.LoadApiMeta, controller.SignUpHandler)
	g.POST("/login", middlewares.LoadApiMeta, controller.LoginHandler)
	g.POST("/auth-token", middlewares.LoadApiMeta, middlewares.JwtAuth(), func(c *gin.Context) {
		c.String(http.StatusOK, "ok userID:%v", c.Value(controller.UserIDKey))
	})
	return r
}
