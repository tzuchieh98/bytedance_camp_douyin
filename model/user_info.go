package model

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Username string `json:"username" gorm:"primary_key"` // 根据 username确定
	Name     string `json:"name"`                        // 默认设置为username
	//FollowCount   int64  `json:"follow_count"`                // 关注数
	//FollowerCount int64  `json:"follower_count"`              // 粉丝数
}

func (UserInfo) TableName() string {
	return "user_infos"
}
