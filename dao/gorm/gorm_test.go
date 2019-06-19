package gorm

import (
	"fmt"
	"github.com/kuhufu/revdol/dao"
	"github.com/kuhufu/revdol/model"
	"log"
	"testing"
)

func TestGetAccountByEmail(t *testing.T) {
	account, err := GetAccountByEmail("huboemail@yeah.net")
	if err != nil {
		log.Println(err)
	}
	log.Println(account)
}

func TestGetAccountById(t *testing.T) {
	account, err := GetAccountById(1)
	if err != nil {
		log.Println(err)
	}
	log.Println(account)
}

func TestRemoveAccount(t *testing.T) {
	a := &model.Account{}
	a.ID = 1
	RemoveAccount(a)
	if a, err := GetAccountById(1); err == nil {
		t.Fatal(a)
	}
}

func TestGetAccountByEmail2(t *testing.T) {
	a := &model.Account{Username: "kuhufu"}
	db.First(&a)
	fmt.Println(a)
}

func TestGetAccountByEmail3(t *testing.T) {

}

func TestRegister(t *testing.T) {
	account, err := dao.Register(&model.Account{Username: "user1"})
	if err != nil {
		log.Println(err)
	}

	log.Println(account)
}

func TestLogin(t *testing.T) {

	account, err := dao.Login("kuhufu", "1111")
	if err != nil {
		log.Println(err)
	}
	log.Println(account)
}

func TestChangePassword(t *testing.T) {
	account, _ := dao.ChangePassword(1, "2222")
	log.Println(account)
}
