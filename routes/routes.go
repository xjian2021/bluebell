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
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.LoadApiMeta)
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
	g.POST("/sign-up", controller.SignUpHandler)
	g.POST("/login", controller.LoginHandler)

	g.Use(middlewares.JwtAuth())
	{
		g.GET("/community", controller.CommunityHandler)
	}
	return r
}
