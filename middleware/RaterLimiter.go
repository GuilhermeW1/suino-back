package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimitMiddleware() gin.HandlerFunc {
	limiter := rate.NewLimiter(10, 20)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(429, gin.H{"error": "muitas requisições, tente novamente"})
			return
		}
		c.Next()
	}
}
