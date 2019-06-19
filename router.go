package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kuhufu/revdol/handler"
	"github.com/kuhufu/revdol/middleware/auth"
)

func router(r *gin.Engine) {

	r.Use(Gzip())
	r.Use(Cors())

	r.Static("/static", "static/")
	r.GET("/", handler.ServeFile("static/index.html"))
	r.Any("/v1/*path", handler.Relay)

	account := r.Group("/account")
	{
		account.GET("/register", handler.ServeFile("static/register.html"))
		account.POST("/register", handler.Register)

		account.GET("/login", handler.ServeFile("static/login.html"))
		account.POST("/login", handler.Login)
		account.GET("/logout", auth.Middleware, handler.Logout)

		account.GET("/changepwd", auth.Middleware, handler.ServeFile("static/changepwd.html"))
		account.POST("/changepwd", auth.Middleware, handler.ChangePassword)

		account.GET("/info", handler.AccountInfo)
	}

	v2 := r.Group("/v2")
	v2.Use(Secure(), auth.Middleware)
	{
		forum := v2.Group("/forum")
		{
			forum.GET("", handler.AllForum)
			forum.GET("/detail/:id", handler.ForumDetail)
			forum.GET("/count/:id", handler.ForumCount)
			forum.GET("/count", handler.AllIdolForumCount)
		}

		idol := v2.Group("/idol")
		{
			idol.GET("/fans-num/:id", handler.FansNum)
			idol.GET("/popular-num/:id", handler.PopularNum)
			idol.GET("/meta/:id", handler.IdolMeta)
			idol.GET("/meta", handler.AllIdolMeta)
			idol.GET("/detail/:id", handler.IdolDetail)
			idol.GET("/detail", handler.IdolList)
		}

		user := v2.Group("/user")
		{
			user.GET("/detail/:id", handler.UserDetail)
			user.GET("/contribute/:id", handler.UserContribute)
		}
	}

	v3 := r.Group("/v3")
	v3.Use(auth.Middleware, CacheControl())
	{
		v3.GET("/forum/detail/:id", handler.ForumDetail)
	}
}
