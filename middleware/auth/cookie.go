package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/kuhufu/revdol/dao/gorm"
	"time"
)

// Hash keys should be at least 32 bytes long
var hashKey = []byte("very-secret")

// Block keys should be 16 bytes (AES-128) or 32 bytes (AES-256) long.
// Shorter keys may weaken the encryption used.
var blockKey = []byte("a-lot-secret")
var s = securecookie.New(hashKey, blockKey)

func Auth(c *gin.Context) {
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
