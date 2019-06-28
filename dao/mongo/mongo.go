package mongo

import "github.com/kuhufu/revdol/dao/Interface"

type Source int

var (
	_ Interface.Revdol = (Source)(nil)
)

func New() Source {
	return Source(0)
}

func (Source) GetAllForum(params map[string]interface{}) interface{} {
	return GetAllForum(params)
}

func (Source) GetUserForumCount(id, currentPage int) interface{} {
	return GetUserForumCount(id, currentPage)
}

func (Source) SearchUser(keyWord string) interface{} {
	return SearchUser(keyWord)
}

func (Source) GetIdolForumCount(id int, page int) interface{} {
	return GetIdolForumCount(id, page)
}

func (Source) GetAllIdolForumCount(page int) interface{} {
	return GetAllIdolForumCount(page)
}
func (Source) GetFansNumById(id int, currentPage int) interface{} {
	return GetFansNumById(id, currentPage)
}

func (Source) GetPopularNumById(id int, currentPage int) interface{} {
	return GetPopularNumById(id, currentPage)
}

func (Source) GetAllIdolMeta(currentPage int) interface{} {
	return GetAllIdolMeta(currentPage)
}

func (Source) GetIdolMetaById(id, page int) interface{} {
	return GetIdolMetaById(id, page)
}

func (Source) GetForumById(id int) interface{} {
	return GetForumById(id)
}

func (Source) GetAllUser(currentPage int) interface{} {
	return GetAllUser(currentPage)
}

func (Source) GetUserById(id int) interface{} {
	return GetUserById(id)
}

func (Source) GetUserContributeById(id int) interface{} {
	return GetUserContributeById(id)
}

func (Source) GetAllUserContribute(currentPage int) interface{} {
	return GetAllUserContribute(currentPage)
}

func (Source) GetIdolById(id int) interface{} {
	return GetIdolById(id)
}

func (Source) GetAllIdol() interface{} {
	return GetAllIdol()
}
