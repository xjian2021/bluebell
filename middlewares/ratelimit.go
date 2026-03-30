package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimit(fillInterval time.Duration, capacity int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, capacity)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) > 0 {
			return
		}
		c.String(http.StatusOK, "寄~~~~")
		c.Abort()
	}
}
