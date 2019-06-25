package casbin

import (
	"github.com/gin-gonic/gin"
)

func Casbin(c *gin.Context) {
	if !Check(c) {
		c.AbortWithStatusJSON(403, gin.H{
			"error": "你没有权限访问",
		})
		return
	}
	c.Next()
}
