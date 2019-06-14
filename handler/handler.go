package handler

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"revdol/dao"
)

const HostURL = "https://starmicro.happyelements.cn"

var urls = map[string]bool{
	"/v1/person/contribution-rank": true,
}

func Relay(c *gin.Context) {
	url := HostURL + c.Request.RequestURI
	//c.Redirect(http.StatusTemporaryRedirect, url)
	req, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
	if err != nil {
		c.Abort()
		return
	}
	req.Header = c.Request.Header
	// add my auth info
	if _, ok := urls[c.Request.URL.Path]; ok {
		req.Header.Set("authorization", "Miinno o44kJAQ3S_nAcl3uSe9GsD8_u7dV6mvT_1548283248")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		c.Abort()
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Abort()
		return
	}
	contentType := resp.Header.Get("Content-Type")
	c.Data(resp.StatusCode, contentType, data)
}

// @Summary forum count
// @Description get string by ID
// @Tags forums
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "Idol ID"
// @Param	page	query	int     true        "page number"
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Router /forum/count/{id} [get]
func ForumCount(c *gin.Context) {
	id := ParamId(c)
	page := QueryPage(c)
	result := dao.GetForumCount(id, page)
	c.JSON(200, result)
}


// @Summary forum count
// @Description get string by ID
// @Tags forums
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "Idol ID"
// @Param	page	query	int     true        "page number"
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Router /forum/count/{id} [get]
func AllIdolForumCount(c *gin.Context) {
	page := QueryPage(c)
	result := dao.GetAllIdolForumCount(page)
	c.JSON(200, result)
}

// @Summary forum count
// @Description get string by ID
// @Tags idols
// @Accept  json
// @Produce  json
// @Param   id	path	int	true	"Idol ID"
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Router /idol/fans-num/{id} [get]
func FansNum(c *gin.Context) {
	id := ParamId(c)
	page := QueryPage(c)
	result := dao.GetFansNumById(id, page)
	c.JSON(200, result)

}

// @Summary forum count
// @Description get string by ID
// @Tags idols
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "Idol ID"
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Router /idol/popular-num/{id} [get]
func PopularNum(c *gin.Context) {
	id := ParamId(c)
	page := QueryPage(c)
	result := dao.GetPopularNumById(id, page)
	c.JSON(200, result)
}

// @Summary all idol meta
// @Description get meta of all idols
// @Tags idols
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Router /idol/meta [get]
func AllIdolMeta(c *gin.Context) {
	page := QueryPage(c)
	result := dao.GetAllIdolMeta(page)
	c.JSON(200, result)
}

// @Summary idol meta
// @Description get string by ID
// @Tags idols
// @Accept  json
// @Produce  json
// @Param	id		path	int     true        "Idol ID"
// @Param	page	query	int     true        "page number"
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Router /idol/meta/{id} [get]
func IdolMeta(c *gin.Context) {
	id := ParamId(c)
	page := QueryPage(c)
	result := dao.GetIdolMetaById(id, page)
	c.JSON(200, result)
}

// @Summary Add a new pet to the store
// @Description get string by ID
// @Tags forums
// @Accept  json
// @Produce  json
// @Param	page	query	int     true        "page number"
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Router /forum [get]
func AllForum(c *gin.Context) {
	page := QueryPage(c)
	result := dao.GetAllForum(page)
	c.JSON(200, result)
}

// @Summary Add a new pet to the store
// @Description get string by ID
// @Tags forums
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "Forum ID"
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Router /forum/detail/{id} [get]
func ForumDetail(c *gin.Context) {
	id := ParamId(c)
	result := dao.GetForumById(id)
	c.JSON(200, result)
}

// @Summary user detail
// @Description get string by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "user ID"
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Router /user/detail/{id} [get]
func UserDetail(c *gin.Context) {
	id := ParamId(c)
	result := dao.GetUserById(id)
	c.JSON(200, result)
}

// @Summary user contribute
// @Description get string by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "user ID"
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Router /user/contribute/{id} [get]
func UserContribute(c *gin.Context) {
	id := ParamId(c)
	result := dao.GetUserContributeById(id)
	c.JSON(200, result)
}

// @Summary idol detail
// @Description idol detail
// @Tags idols
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "Idol ID"
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Router /idol/detail/{id} [get]
func IdolDetail(c *gin.Context) {
	id := ParamId(c)
	result := dao.GetIdolById(id)
	c.JSON(200, result)
}

// @Summary all idol detail
// @Description detail of all idols
// @Tags idols
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Security ApiKeyAuth
// @Router /idol/detail [get]
func IdolList(c *gin.Context) {
	result := dao.GetAllIdol()
	c.JSON(200, result)
}
