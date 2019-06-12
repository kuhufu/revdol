package dao

import "revdol/model"

type Account interface {
	Register(account *model.Account) (*model.Account, error)
	Login(identity, password string) (*model.Account, error)
	ChangePassword(id uint, newPwd string) (*model.Account, error)
}

type Source interface {
	Account
	Interface
}

type Interface interface {
	GetForumCount(id string) interface{}
	GetFansNumById(id string) interface{}
	GetPopularNumById(id string) interface{}
	GetAllIdolMeta() interface{}
	GetIdolMetaById(id string) interface{}
	GetForumById(id string) []byte
	GetAllForum() []string
	GetAllUser() []string
	GetUserById(id string) []byte
	GetUserContributeById(id string) []byte
	GetAllUserContribute() func() []string
	GetIdolById(id string) []byte
	GetAllIdol() []byte
}
