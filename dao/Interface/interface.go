package Interface

import "github.com/kuhufu/revdol/model"

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
	GetAllIdolMeta(currentPage int) interface{}
	GetIdolMetaById(id, currentPage int) interface{}
	GetFansNumById(id, currentPage int) interface{}
	GetPopularNumById(id, currentPage int) interface{}

	GetAllForum(params map[string]interface{}) interface{}
	GetForumById(id int) interface{}

	GetAllUser(currentPage int) interface{}
	GetUserById(id int) interface{}
	GetUserContributeById(id int) interface{}
	GetAllUserContribute(currentPage int) interface{}

	GetIdolById(id int) interface{}
	GetAllIdol() interface{}

	GetIdolForumCount(id, currentPage int) interface{}
	GetAllIdolForumCount(currentPage int) interface{}
	GetUserForumCount(id, currentPage int) interface{}

	SearchUser(keyWord string, currentPage int) interface{}
	SearchForum(filed, keyWord string, currentPage int) interface{}
}
