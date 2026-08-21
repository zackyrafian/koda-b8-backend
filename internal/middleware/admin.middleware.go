package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "ADMIN" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "forbidden: admin access required",
			})
			return
		}
		c.Next()
	}
}
