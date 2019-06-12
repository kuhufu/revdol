package dao

import (
	"revdol/dao/redisSource"
)

var dbSource Interface

func init() {
	dbSource = redisSource.DB
}

func GetForumCount(id string) interface{} {
	return dbSource.GetForumCount(id)
}

func GetFansNumById(id string) interface{} {
	return dbSource.GetFansNumById(id)
}

func GetPopularNumById(id string) interface{} {
	return dbSource.GetPopularNumById(id)
}

func GetAllIdolMeta() interface{} {
	return dbSource.GetAllIdolMeta()
}

func GetIdolMetaById(id string) interface{} {
	return dbSource.GetIdolMetaById(id)
}

func GetForumById(id string) []byte {
	return dbSource.GetForumById(id)
}

func GetAllForum() []string {
	return dbSource.GetAllForum()
}

func GetAllUser() []string {
	return dbSource.GetAllUser()
}

func GetUserById(id string) []byte {
	return dbSource.GetUserById(id)
}

func GetUserContributeById(id string) []byte {
	return dbSource.GetUserContributeById(id)
}

func GetAllUserContribute() func() []string {
	return dbSource.GetAllUserContribute()
}

func GetIdolById(id string) []byte {
	return dbSource.GetIdolById(id)
}

func GetAllIdol() []byte {
	return dbSource.GetAllIdol()
}
