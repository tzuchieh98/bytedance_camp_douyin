package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	UserID      int64  `json:"user_info_id"`
	ToUserID    int64  `json:"video_id"`
	Content     string `json:"content"`
	PublishDate string `json:"publish_date"`
}

func (Message) TableName() string {
	return "messages"
}
