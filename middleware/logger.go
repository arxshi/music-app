package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime)
		log.Printf("[INFO] %s %s %s %v", c.Request.Method, c.Request.URL.Path, c.ClientIP(), duration)
	}
}
