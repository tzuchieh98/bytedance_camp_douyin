package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserInfoID  int64  `json:"user_info_id"`
	VideoID     int64  `json:"video_id"`
	Content     string `json:"content"`
	PublishDate string `json:"publish_date"`
}

func (Comment) TableName() string {
	return "comments"
}
