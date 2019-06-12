package dao

import (
	"errors"
	"revdol/dao/gormSource"
	"revdol/model"
)

const (
	User  = "user"
	Admin = "admin"
	Root  = "root"
)

func Register(account *model.Account) (*model.Account, error) {
	pwd := ""
	pwd, account.Password = account.Password, pwd

	count := 0
	if gormSource.DB.Model(account).Where(account).Count(&count); count != 0 {
		return nil, errors.New("identity already exist")
	}

	account.Role = User
	account.Password = pwd
	gormSource.DB.Create(account)
	return account, nil
}

func Login(identity, password string) (*model.Account, error) {
	a := &model.Account{}
	count := 0
	gormSource.DB.Where("(username = ? or email = ?) AND password = ?", identity, identity, password).First(a).Count(&count)

	if count == 0{
		return nil, errors.New("wrong identity or password")
	}

	return a, nil
}

func ChangePassword(id uint, newPwd string) (*model.Account, error) {
	a := &model.Account{}
	a.ID = id
	gormSource.DB.Model(a).UpdateColumns(model.Account{Password: newPwd}).First(a)
	return a, nil
}

func AccountExists(account *model.Account)  {

}