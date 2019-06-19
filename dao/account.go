package dao

import (
	"github.com/kuhufu/revdol/dao/gorm"
	"github.com/kuhufu/revdol/model"
)

func Register(account *model.Account) (*model.Account, error) {
	return gorm.Register(account)
}

func Login(identity, password string) (*model.Account, error) {
	return gorm.Login(identity, password)
}

func ChangePassword(id uint, newPwd string) (*model.Account, error) {
	return gorm.ChangePassword(id, newPwd)
}

func AccountExists(account *model.Account) bool {
	return gorm.AccountExists(account)
}
