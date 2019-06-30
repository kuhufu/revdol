package redis

import "github.com/kuhufu/revdol/dao/Interface"

type Source int

var (
	_ Interface.Revdol = Source(0)
)

func New() Source {
	return Source(0)
}

func (s Source) GetAllIdolMeta(currentPage int) interface{} {
	return GetAllIdolMeta()
}

func (s Source) GetIdolMetaById(id, currentPage int) interface{} {
	return GetIdolMetaById(string(id))
}

func (s Source) GetFansNumById(id, currentPage int) interface{} {
	return GetFansNumById(string(id))
}

func (s Source) GetPopularNumById(id, currentPage int) interface{} {
	return GetPopularNumById(string(id))
}

func (s Source) GetAllForum(params map[string]interface{}) interface{} {
	panic("implement me")
}

func (s Source) GetForumById(id int) interface{} {
	return GetForumById(string(id))
}

func (s Source) GetAllUser(currentPage int) interface{} {
	panic("implement me")
}

func (s Source) GetUserById(id int) interface{} {
	return GetUserById(string(id))
}

func (s Source) GetUserContributeById(id int) interface{} {
	return GetUserContributeById(string(id))
}

func (s Source) GetAllUserContribute(currentPage int) interface{} {
	panic("implement me")
}

func (s Source) GetIdolById(id int) interface{} {
	return GetIdolById(string(id))
}

func (s Source) GetAllIdol() interface{} {
	return GetAllIdol()
}

func (s Source) GetIdolForumCount(id, currentPage int) interface{} {
	return GetIdolForumCount(string(id))
}

func (s Source) GetAllIdolForumCount(currentPage int) interface{} {
	panic("implement me")
}

func (s Source) GetUserForumCount(id, currentPage int) interface{} {
	panic("implement me")
}

func (s Source) SearchUser(keyWord string) interface{} {
	panic("implement me")
}
