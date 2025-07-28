package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()

		fmt.Printf("[LOG] %s %s | %d | %s\n", c.Request.Method, c.Request.URL.Path, status, latency)
	}
}
