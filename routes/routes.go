package routes

import (
	"net/http"
	"time"

	"github.com/gin-contrib/pprof"
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
	r.GET("/ping", middlewares.RateLimit(2*time.Second, 1), func(c *gin.Context) {
		c.String(http.StatusOK, "ok %s", time.Now().String())
	})
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
		g.GET("/community/:id", controller.CommunityDetailHandler)
		g.POST("/post", controller.CreatePostHandler)
		g.GET("/post/:id", controller.PostDetailHandler)
		g.GET("/posts", controller.PostListHandler)
		g.POST("/vote", controller.VoteHandler)
	}
	pprof.Register(r)
	return r
}
