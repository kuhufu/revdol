package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"log"
	. "revdol/config"
)

func Session() gin.HandlerFunc {
	//store := cookie.NewStore([]byte("secret"))
	store, _ := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret"))

	store.Options(sessions.Options{
		MaxAge: 60,
	})
	return sessions.Sessions("revdol", store)
}

func Secure() gin.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny:     true,
		SSLRedirect:   true,
		SSLHost:       "localhost" + Config.Https_port,
		IsDevelopment: Config.Dev, //开发模式下，这些设置都会被忽略
	})

	return func(c *gin.Context) {
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			log.Println(err)
			c.Abort()
			return
		}
		c.Next()
	}
}

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AddAllowHeaders("authorization")
	config.AllowAllOrigins = true
	return cors.New(config)
}

func CacheControl() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "max-age=31536000")
	}
}

func Gzip() gin.HandlerFunc {
	return gzip.Gzip(gzip.DefaultCompression)
}
