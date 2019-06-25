package handler

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/kuhufu/revdol/constant"
	"github.com/kuhufu/revdol/dao"
	"github.com/kuhufu/revdol/middleware/auth"
	"github.com/kuhufu/revdol/model"
	"net/http"
	"time"
)

type LoginForum struct {
	Identity string `form:"identity" json:"identity" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type RegisterForum struct {
	Username string `form:"username" json:"username" binding:"required"`
	Email    string `form:"email" json:"email" binding:"omitempty,email"`
	Password string `form:"password" json:"password" binding:"required,min=4,printascii"`
}

type ChangePwdForm struct {
	Password        string `form:"password" json:"password" binding:"min=4,required,printascii,eqfield=ConfirmPassword"`
	ConfirmPassword string `form:"confirmPassword" json:"confirmPassword" binding:"min=4,required,printascii"`
}

func Login(c *gin.Context) {
	login := LoginForum{}
	if err := c.Bind(&login); err != nil {
		return
	}
	account, err := dao.Login(login.Identity, login.Password)
	c.Set(constant.LoginInfo, gin.H{
		constant.AccountKey: account,
		"error":             err,
	})
	auth.LoginHandler(c)
}

func Logout(c *gin.Context) {
	if !auth.Mw.SendCookie {
		auth.Mw.Unauthorized(c, 403, jwt.ErrForbidden.Error())
		return
	}
	cookie, err := c.Request.Cookie(auth.Mw.CookieName)
	if err != nil {
		auth.Mw.Unauthorized(c, 403, jwt.ErrEmptyCookieToken.Error())
		return
	}
	cookie.MaxAge = -1
	cookie.Expires = time.Now()
	cookie.Value = ""
	cookie.Path = "/"
	c.Header("Set-Cookie", cookie.String())

	c.JSON(200, gin.H{
		"status": "success",
		"action": "logout",
	})
}

func Register(c *gin.Context) {
	f := &RegisterForum{}
	if err := c.ShouldBind(f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	a, err := dao.Register(&model.Account{
		Username: f.Username,
		Email:    f.Email,
		Password: f.Password,
	})

	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, a)
}

func ChangePassword(c *gin.Context) {
	f := &ChangePwdForm{}
	if err := c.ShouldBind(f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	account, _ := c.Get(constant.AccountKey)
	dao.ChangePassword(account.(*model.Account).ID, f.Password)
	c.JSON(http.StatusOK, gin.H{
		"message": "change password successful",
	})
}

func AccountInfo(c *gin.Context) {
	a, _ := c.Get(constant.AccountKey)
	c.JSON(http.StatusOK, a)
}
