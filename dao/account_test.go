package dao

import (
	"log"
	"revdol/model"
	"testing"
)


func TestRegister(t *testing.T) {
	account, err := Register(&model.Account{Username:"user1"})
	if err != nil {
		log.Println(err)
	}

	log.Println(account)
}

func TestLogin(t *testing.T) {

	account, err := Login("kuhufu", "1111")
	if err != nil {
		log.Println(err)
	}
	log.Println(account)
}

func TestChangePassword(t *testing.T) {
	account, _ := ChangePassword(1, "2222")
	log.Println(account)
}


