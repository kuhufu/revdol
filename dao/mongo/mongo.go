package mongo


type MdbSource int

func New() MdbSource {
	return MdbSource(0)
}

func (MdbSource) GetForumCount(id int, page int) interface{} {
	return GetForumCount(id,  page)
}
func (MdbSource) GetAllIdolForumCount(page int) interface{} {
	return GetAllIdolForumCount(page)
}

func (MdbSource) GetFansNumById(id int, currentPage int) interface{} {
	return GetFansNumById(id, currentPage)
}

func (MdbSource) GetPopularNumById(id int, currentPage int) interface{} {
	return GetPopularNumById(id, currentPage)
}

func (MdbSource) GetAllIdolMeta(currentPage int) interface{} {
	return GetAllIdolMeta(currentPage)
}

func (MdbSource) GetIdolMetaById(id, page int) interface{} {
	return GetIdolMetaById(id, page)
}

func (MdbSource) GetForumById(id int) interface{} {
	return GetForumById(id)
}

func (MdbSource) GetAllForum(currentPage int) interface{} {
	return GetAllForum(currentPage)
}

func (MdbSource) GetAllUser(currentPage int) interface{}{
	return GetAllUser(currentPage)
}

func (MdbSource) GetUserById(id int) interface{} {
	return GetUserById(id)
}

func (MdbSource) GetUserContributeById(id int) interface{} {
	return GetUserContributeById(id)
}

func (MdbSource) GetAllUserContribute(currentPage int) interface{} {
	return GetAllUserContribute(currentPage)
}

func (MdbSource) GetIdolById(id int) interface{} {
	return GetIdolById(id)
}

func (MdbSource) GetAllIdol() interface{} {
	return GetAllIdol()
}
