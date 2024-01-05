package middlewares

import (
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

// NewTimeOutMiddleware increases the timeout of the gin.HandlerFunc it is applied to
func NewTimeOutMiddleware(seconds time.Duration) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(seconds),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
	)
}
