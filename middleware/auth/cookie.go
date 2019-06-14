package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"revdol/dao/gorm"
	"time"
)

func Auth(c *gin.Context)  {
	if c.Request.URL.Path == "/login" {
		c.Next()
		return
	}

	tokenSession := sessions.Default(c)
	uid := tokenSession.Get("uid")
	expires := tokenSession.Get("expires")
	if uid == nil || uid.(uint32) == 0 {
		c.AbortWithStatusJSON(403, gin.H{
			"error": "no authorization",
		})
		return
	}

	expireTime := expires.(int64)
	if expireTime < time.Now().Unix() {
		c.AbortWithStatusJSON(403, gin.H{
			"error": "no authorization",
		})
		return
	}

	account, _ := gorm.GetAccountById(uid.(uint))
	c.Set("user", account)
	c.Next()
}
