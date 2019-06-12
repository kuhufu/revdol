package model

import "time"

type Forum struct {
	JsonContent string `gorm:"type:json" json:"-""`

	ID             uint      `json:"id" gorm:"primary_key"`
	UserID         uint      `json:"user_id"`
	IdolID         uint      `json:"idol_id"`
	Title          string    `json:"title"`
	Content        string    `json:"content" gorm:"type:text"`
	CommentsCounts int       `json:"comment_counts"`
	IsOriginal     string    `json:"is_original"`
	IsPick         string    `json:"is_pick"`
	IsTop          string    `json:"is_top"`
	PraiseCounts   int       `json:"praise_counts"`
	Status         int       `json:"status"`
	Tag            int       `json:"tag"`
	Type           int       `json:"type"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}
