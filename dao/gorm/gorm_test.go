package gormSource

import (
	"fmt"
	"log"
	"revdol/model"
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
	RemoveAccount(&model.Account{ID: 1})
	if a, err := GetAccountById(1); err == nil {
		t.Fatal(a)
	}
}

func TestGetAccountByEmail2(t *testing.T) {
	a := &model.Account{Username:"kuhufu"}
	DB.First(&a)
	fmt.Println(a)
}

func TestGetAccountByEmail3(t *testing.T) {

}
