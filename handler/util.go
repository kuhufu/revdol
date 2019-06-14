package handler

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func QueryPage(c *gin.Context) int {
	p := c.Query("page")
	page, _ := strconv.Atoi(p)
	if page == 0 {
		page = 1
	}
	return page
}

func ParamId(c *gin.Context) int {
	i := c.Param("id")
	id, _ := strconv.Atoi(i)
	return id
}
