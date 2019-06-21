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

func ForumQueryParams(c *gin.Context) (result map[string]interface{}, err error) {
	result = map[string]interface{}{}
	params := []string{"user_id", "idol_id", "page"}
	for _, param := range params {
		if value, exist := c.GetQuery(param); exist {
			v, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			result[param] = v
		}
	}
	if _, ok := result["page"]; !ok {
		result["page"] = 1
	}
	return
}
