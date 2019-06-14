package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"log"
)

func ByteSession(sname string) gin.HandlerFunc {
	store := cookie.NewStore([]byte("byte"))
	store.Options(sessions.Options{
		Domain: "/",
		MaxAge: 30 * 60,
	})
	return sessions.Sessions(sname, store)
}

func RedisSession(sname string) gin.HandlerFunc {
	store, err := redis.NewStore(10, "tcp", "127.0.0.1:7001", "", []byte("redis"))
	if err != nil {
		log.Fatal(err)
	}

	store.Options(sessions.Options{
		Domain: "/",
		MaxAge: 30 * 60,
	})
	return sessions.Sessions(sname, store)
}