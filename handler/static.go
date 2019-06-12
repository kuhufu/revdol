package handler

import "github.com/gin-gonic/gin"

func ServeFile(path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.File(path)
	}
}
