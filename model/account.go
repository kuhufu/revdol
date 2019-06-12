package model

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	Password    string `json:"-"`
	Username    string
	Nickname    string
	Role        string `gorm:"type:varchar(20);not nul"`
	Email       string `gorm:"type:varchar(100);index;not null"`
	PhoneNumber string `gorm:"type:varchar(20)"`

	RevdolUserID uint
}
