package dao

import (
	"github.com/kuhufu/revdol/dao/Interface"
	"github.com/kuhufu/revdol/dao/mongo"
)

var source Interface.Revdol = mongo.New()

func GetForumCount(id int, page int) interface{} {
	return source.GetForumCount(id, page)
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

func GetAllForum(currentPage int) interface{} {
	return source.GetAllForum(currentPage)
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
