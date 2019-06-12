package model

import "time"

type User struct {
	JsonContent string `gorm:"type:json"`

	ID                uint      `json:"id" gorm:"primary_key"`
	Openid            string    `json:"openid"`
	Uid               int       `json:"uid"`
	Sex               int       `json:"sex"`
	Nickname          string    `json:"nickname"`
	Birth             string    `json:"birth"`
	Tel               string    `json:"tel"`
	Headimg           string    `json:"headimg"`
	Address           string    `json:"address"`
	Skin              string    `json:"skin"`
	Status            int       `json:"status"`
	ExcuseTime        int       `json:"excuse_time"`
	Password          string    `json:"password"`
	AreaCode          string    `json:"area_code"`
	Invitation        string    `json:"invitation"`
	SourceFlag        string    `json:"source_flag"`
	CreatedAt         time.Time `json:"-"` // 可能换成 unix时间
	UpdatedAt         time.Time `json:"-"`
	//
	//Forums      []Forum
	//Contributes []Contribute
}
