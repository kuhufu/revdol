package Interface

import "revdol/model"

type Account interface {
	Register(account *model.Account) (*model.Account, error)
	Login(identity, password string) (*model.Account, error)
	ChangePassword(id uint, newPwd string) (*model.Account, error)
}

type Source interface {
	Account
	Revdol
}

type Revdol interface {
	GetForumCount(id, currentPage int) interface{}
	GetAllIdolForumCount(currentPage int) interface{}
	GetFansNumById(id, currentPage int) interface{}
	GetPopularNumById(id, currentPage int) interface{}
	GetAllIdolMeta(currentPage int) interface{}
	GetIdolMetaById(id, currentPage int) interface{}
	GetForumById(id int) interface{}
	GetAllForum(currentPage int) interface{}
	GetAllUser(currentPage int) interface{}
	GetUserById(id int) interface{}
	GetUserContributeById(id int) interface{}
	GetAllUserContribute(currentPage int)interface{}
	GetIdolById(id int) interface{}
	GetAllIdol() interface{}
}
