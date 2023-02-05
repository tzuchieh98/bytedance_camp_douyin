package model

import "gorm.io/gorm"

type UserLogin struct {
	gorm.Model
	Username string `json:"username" gorm:"primary_key"` // 登录用户名，通过 username 查找用户信息
	Password string `json:"password"`                    // 登录密码
}

func (UserLogin) TableName() string {
	return "user_logins"
}
