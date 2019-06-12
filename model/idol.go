package model

import "time"

type Idol struct {
	JsonContent string `gormSource:"type:json"`

	ID                  uint      `json:"id" gormSource:"primary_key"`
	PopularNum          int       `json:"popular_num"`
	FansNum             int       `json:"fans_num"`
	AttentionNum        int       `json:"attention_num"`
	Fullname            string    `json:"fullname"`
	Nickname            string    `json:"nickname"`
	NationalRepresent   string    `json:"national_represent"`
	Headimg             string    `json:"headimg"`
	ShortDesc           string    `json:"short_desc"`
	Desc                string    `json:"desc"`
	Poster              string    `json:"poster"`
	MainPoster          string    `json:"main_poster"`
	Height              int       `json:"height"`
	Age                 int       `json:"age"`
	Like                string    `json:"like"`
	Birth               string    `json:"birth"`
	RepresentFlower     string    `json:"represent_flower"`
	Status              int       `json:"status"`
	DarkSkin            string    `json:"dark_skin"`
	LightSkin           string    `json:"light_skin"`
	PV                  string    `json:"pv"`
	Voice               string    `json:"voice"`
	Bg                  string    `json:"bg"`
	IntroduceBackground string    `json:"introduce_background"`
	IntroduceProspect   string    `json:"introduce_prospect"`
	AssistTop           string    `json:"assist_top"`
	BloodType           string    `json:"blood_type"`
	Birthplace          string    `json:"birthplace"`
	Likes               string    `json:"likes"`
	Hates               string    `json:"hates"`
	IntroduceMiddle     string    `json:"introduce_middle"`
	UpdatedAt           time.Time `json:"-"`
}
