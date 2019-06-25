package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	. "github.com/kuhufu/revdol/config"
	"github.com/kuhufu/revdol/constant"
	"github.com/kuhufu/revdol/dao/gorm"
	"github.com/kuhufu/revdol/middleware/casbin"
	"github.com/kuhufu/revdol/model"
	"net/http"
	"sync"
	"time"
)

const identityKey = "uid"

var LoginHandler = Mw.LoginHandler

var CookieName = Mw.CookieName

var cache = sync.Map{}

var Middleware = func(c *gin.Context) {
	//是否启用 auth（认证和授权）
	if !Config.EnableAuth {
		c.Next()
		return
	}
	Mw.MiddlewareFunc()(c)
}

var Mw, _ = jwt.New(&jwt.GinJWTMiddleware{
	Realm:          "test zone",
	Key:            []byte("secret key"),
	Timeout:        2 * time.Hour,
	MaxRefresh:     time.Minute,
	IdentityKey:    identityKey,
	SendCookie:     true,
	SecureCookie:   Config.Dev,
	CookieHTTPOnly: true,

	// 1. 登录时会调用，返回的是可以添加到 payload 数据
	Authenticator: func(c *gin.Context) (interface{}, error) {
		m, exists := c.Get(constant.LoginInfo)
		if !exists {
			return nil, nil
		}
		loginInfo := m.(gin.H)
		if err := loginInfo["error"]; err != nil {
			return nil, err.(error)
		}
		return loginInfo[constant.AccountKey], nil
	},

	// 2. Authenticator 生成的 data 回传入 PayloadFunc 经过一些处理后生成 jwt.MapClaims（map[string]interface{}）
	// PayloadFunc 调用完成后会生成 token 返回给客户端
	PayloadFunc: func(data interface{}) jwt.MapClaims {
		if account, ok := data.(*model.Account); ok {
			return jwt.MapClaims{
				identityKey: account.ID, // ID 就够了
			}
		}
		return jwt.MapClaims{}
	},

	// 3. 客户端带着token发送请求后，服务器会检查 token 是否过期。
	// 没有过期 则调用 IdentityHandler 解析 payload 生成 jwt.MapClaims类型数据，
	// 		在IdentityHandler经过某些处理后，将IdentityHandler返回的数据传给 Authorizator
	// 如果过期 则调用 Unauthorized 函数
	IdentityHandler: func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return uint(claims[identityKey].(float64))
	},

	// 4. IdentityHandler之后调用
	// true	 表示授权成功，将会调用 router handler
	// false 表示授权失败，将会调用 Unauthorized
	Authorizator: func(identity interface{}, c *gin.Context) bool {
		s, _ := c.Cookie(CookieName)
		account, ok := cache.Load(s)
		if !ok {
			a, err := gorm.GetAccountById(identity.(uint))
			if err != nil {
				return false
			}
			account = a
			s, _ := c.Cookie(CookieName)
			cache.Store(s, account) //加入缓存,避免对同一个 jwt token 的多次 decode 和请求 mysql查 询 account
		}
		c.Set(constant.AccountKey, account)
		return casbin.Check(c) //这里使用casbin做权限管理
	},

	// 认证失败后调用
	Unauthorized: func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	},

	//
	LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
		c.JSON(http.StatusOK, gin.H{
			"name":   "Bearer",
			"code":   http.StatusOK,
			"token":  token,
			"expire": expire.Format(time.RFC3339),
		})
	},

	TokenLookup: "cookie: jwt, header: Authorization, query: token",

	TokenHeadName: "Bearer",
})
