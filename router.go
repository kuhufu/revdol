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

		account.GET("/info", auth.Middleware, handler.AccountInfo)
	}

	v2 := r.Group("/v2")
	v2.Use(Secure(), auth.Middleware)
	{
		forum := v2.Group("/forum")
		{
			forum.GET("", handler.AllForum)
			forum.GET("/detail/:id", handler.ForumDetail)
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

		count := v2.Group("/count")
		{
			count.GET("/forum/user/:id", handler.CountUserForum)
			count.GET("/forum/idol", handler.CountAllIdolForum)
			count.GET("/forum/idol/:id", handler.CountIdolForum)
		}

		search := v2.Group("/search")
		{
			search.GET("", handler.ServeFile("static/search.html"))
			search.GET("/user", handler.SearchUser)
		}
	}
}
