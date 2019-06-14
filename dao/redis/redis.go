package redis

type Source int

func New() *Source {
	i := Source(1)
	return &i
}

func (s *Source) GetForumCount(id string) interface{} {
	return GetForumCount(id)
}

func (s *Source) GetFansNumById(id string) interface{} {
	return GetFansNumById(id)
}

func (s *Source) GetPopularNumById(id string) interface{} {
	return GetPopularNumById(id)
}

func (s *Source) GetAllIdolMeta() interface{} {
	return GetAllIdolMeta()
}

func (s *Source) GetIdolMetaById(id string) interface{} {
	return GetIdolMetaById(id)
}

func (s *Source) GetForumById(id string) []byte {
	return GetForumById(id)
}

func (s *Source) GetAllForum() []string {
	return GetAllForum()
}

func (s *Source) GetAllUser() []string {
	return GetAllUser()
}

func (s *Source) GetUserById(id string) []byte {
	return GetUserById(id)
}

func (s *Source) GetUserContributeById(id string) []byte {
	return GetUserContributeById(id)
}

func (s *Source) GetAllUserContribute() func() []string {
	return GetAllUserContribute()
}

func (s *Source) GetIdolById(id string) []byte {
	return GetIdolById(id)
}

func (s *Source) GetAllIdol() []byte {
	return GetAllIdol()
}
