package casbin

import (
	"github.com/gin-gonic/gin"
	"github.com/kuhufu/revdol/model"
)

func Casbin(c *gin.Context) {
	account := &model.Account{}
	if u, exists := c.Get("account"); !exists {
		account.Role = "anonymous"
		account.Username = "anonymous"
	} else {
		account = u.(*model.Account)
	}

	e := GetEnforce()

	role := account.Role
	path := c.Request.URL.Path
	method := c.Request.Method

	if !e.Enforce(role, path, method) {
		c.AbortWithStatusJSON(403, gin.H{
			"error":    "你没有权限访问",
			"username": account.Username,
			"method":   method,
			"path":     path,
		})
		return
	}
	c.Next()
}
