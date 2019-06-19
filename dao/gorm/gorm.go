package gorm

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "github.com/kuhufu/revdol/config"
	"github.com/kuhufu/revdol/model"
	"log"
)

var db *gorm.DB

const (
	User  = "user"
	Admin = "admin"
	Root  = "root"
)

func init() {
	var err error
	db, err = gorm.Open(Config.Gorm.Provider, Config.Gorm.URL)
	if err != nil {
		log.Println(err)
	}

	db.LogMode(Config.Gorm.LogMode)

	db.AutoMigrate(&model.Account{})
	db.AutoMigrate(&model.User{}, &model.Forum{}, &model.Idol{}, &model.Contribute{})
}

func GetAccountByEmail(email string) (*model.Account, error) {
	a := &model.Account{}
	if notFount := db.Where("email = ?", email).First(a).RecordNotFound(); notFount {
		return nil, errors.New("account not exist")
	}
	return a, nil
}

func GetAccountById(id uint) (*model.Account, error) {
	a := &model.Account{}
	a.ID = id

	if notFound := db.First(a).RecordNotFound(); notFound {
		return nil, errors.New("account not exist")
	}
	return a, nil
}

func RemoveAccount(account *model.Account) {
	db.Delete(account)
}

func Register(account *model.Account) (*model.Account, error) {
	pwd := ""
	pwd, account.Password = account.Password, pwd

	count := 0
	if db.Model(account).Where(account).Count(&count); count != 0 {
		return nil, errors.New("identity already exist")
	}

	account.Role = User
	account.Password = pwd
	db.Create(account)
	return account, nil
}

func Login(identity, password string) (*model.Account, error) {
	a := &model.Account{}
	count := 0
	db.Where("(username = ? or email = ?) AND password = ?", identity, identity, password).First(a).Count(&count)

	if count == 0 {
		return nil, errors.New("wrong identity or password")
	}

	return a, nil
}

func ChangePassword(id uint, newPwd string) (*model.Account, error) {
	a := &model.Account{}
	a.ID = id
	db.Model(a).UpdateColumns(model.Account{Password: newPwd}).First(a)
	return a, nil
}

func AccountExists(account *model.Account) bool {
	return false
}
