package model

import "time"

type Contribute struct {
	JsonContent string `gorm:"type:json"`

	ID     uint `json:"id" gorm:"primary_key"`
	UserID uint `json:"user_id"`
	IdolID uint `json:"idol_id"`

	Point     int `json:"point"`
	Level     int `json:"level"`
	Status    int `json:"status"`

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
