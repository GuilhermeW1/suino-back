package middleware

import "github.com/gin-gonic/gin"

func APIKeyMiddleware(validKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-Key")
		if key == "" || key != validKey {
			c.AbortWithStatusJSON(401, gin.H{"error": "não autorizado"})
			return
		}
		c.Next()
	}
}
