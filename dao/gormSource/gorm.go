package gormSource

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	. "revdol/config"
	"revdol/model"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(Config.Gorm.Provider, Config.Gorm.URL)
	//Gorm, err = gorm.Open("mysql", "root:7266@/revdol?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	DB.LogMode(Config.Gorm.Log)

	DB.AutoMigrate(&model.Account{})
	DB.AutoMigrate(&model.User{}, &model.Forum{}, &model.Idol{}, &model.Contribute{})
}

func GetAccountByEmail(email string) (*model.Account, error) {
	a := &model.Account{}
	if notFount := DB.Where("email = ?", email).First(a).RecordNotFound(); notFount {
		return nil, errors.New("account not exist")
	}
	return a,nil
}

func GetAccountById(id uint) (*model.Account, error) {
	a := &model.Account{}
	a.ID = id


	if notFound := DB.First(a).RecordNotFound(); notFound {
		return nil, errors.New("account not exist")
	}
	return a, nil
}

func RemoveAccount(account *model.Account) {
	DB.Delete(account)
}