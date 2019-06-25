package dao

import (
	"encoding/json"
	"github.com/kuhufu/revdol/cache"
	"github.com/kuhufu/revdol/dao/Interface"
	"github.com/kuhufu/revdol/dao/mongo"
)

var source Interface.Revdol = mongo.New()

func GetIdolForumCount(id int, page int) interface{} {
	return source.GetIdolForumCount(id, page)
}

func GetAllIdolForumCount(page int) interface{} {
	return source.GetAllIdolForumCount(page)
}

func GetFansNumById(id, page int) interface{} {
	return source.GetFansNumById(id, page)
}

func GetPopularNumById(id, page int) interface{} {
	return source.GetPopularNumById(id, page)
}

func GetAllIdolMeta(currentPage int) interface{} {
	return source.GetAllIdolMeta(currentPage)
}

func GetIdolMetaById(id, page int) interface{} {
	return source.GetIdolMetaById(id, page)
}

func GetForumById(id int) interface{} {
	return source.GetForumById(id)
}

func GetAllForum(params map[string]interface{}) interface{} {
	return source.GetAllForum(params)
}

func GetUserForumCount(id, currentPage int) interface{} {
	return source.GetUserForumCount(id, currentPage)
}

func GetAllUser(currentPage int) interface{} {
	return source.GetAllUser(currentPage)
}

func GetUserById(id int) interface{} {
	return source.GetUserById(id)
}

func GetUserContributeById(id int) interface{} {
	return source.GetUserContributeById(id)
}

func GetAllUserContribute(currentPage int) interface{} {
	return source.GetAllUserContribute(currentPage)
}

func GetIdolById(id int) interface{} {
	return source.GetIdolById(id)
}

func GetAllIdol() interface{} {
	return source.GetAllIdol()
}

func SearchUser(keyWord string) interface{} {
	cacheKeyWord := "search:wd:" + keyWord
	result, err := cache.GetUnmarshal(cacheKeyWord)
	if err != nil {
		data, _ := json.Marshal(source.SearchUser(keyWord))
		cache.Set(cacheKeyWord, data, 10)
		result, _ := cache.GetUnmarshal(cacheKeyWord)
		return result
	}
	return result
}
