package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"bluebell/controller"
	"bluebell/logger"
	"bluebell/pkg/snowflake"
	"bluebell/settings"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "version:%s", settings.C.Version)
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{"msg_id": snowflake.GenID(), "msg": "path not found"},
		)
	})
	r.POST("/sign-up", controller.SignUpHandler)
	return r
}
