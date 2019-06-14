package mongoSource


type mdbSource int

func (mdbSource) GetForumCount(id string) interface{} {
	return GetForumCount(id)
}

func (mdbSource) GetFansNumById(id string) interface{} {
	panic("implement me")
}

func (mdbSource) GetPopularNumById(id string) interface{} {
	panic("implement me")
}

func (mdbSource) GetAllIdolMeta() interface{} {
	panic("implement me")
}

func (mdbSource) GetIdolMetaById(id string) interface{} {
	panic("implement me")
}

func (mdbSource) GetForumById(id string) []byte {
	panic("implement me")
}

func (mdbSource) GetAllForum() []string {
	panic("implement me")
}

func (mdbSource) GetAllUser() []string {
	panic("implement me")
}

func (mdbSource) GetUserById(id string) []byte {
	panic("implement me")
}

func (mdbSource) GetUserContributeById(id string) []byte {
	panic("implement me")
}

func (mdbSource) GetAllUserContribute() func() []string {
	panic("implement me")
}

func (mdbSource) GetIdolById(id string) []byte {
	panic("implement me")
}

func (mdbSource) GetAllIdol() []byte {
	panic("implement me")
}
