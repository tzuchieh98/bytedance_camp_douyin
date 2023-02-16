package dal

import (
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"github.com/linzijie1998/bytedance_camp_douyin/model"
	"gorm.io/gorm"
)

func CreateUserInfo(info *model.UserInfo) error {
	return global.DOUYIN_DB.Create(info).Error
}

func DeleteUserInfoByUsername(username string) error {
	return global.DOUYIN_DB.Where("username = ?", username).Delete(&model.UserInfo{}).Error
}

func UpdateUserInfo(info *model.UserInfo) error {
	return global.DOUYIN_DB.Updates(info).Error
}

func QueryUserInfoByUsername(username string) (userInfos []model.UserInfo, err error) {
	err = global.DOUYIN_DB.Find(&userInfos, "username = ?", username).Error
	return
}

func QueryUserInfoByUserID(id int64) (userInfos []model.UserInfo, err error) {
	err = global.DOUYIN_DB.Find(&userInfos, "id = ?", id).Error
	return
}

// UserRegister 注册: 同时向数据库存储登录信息UserLogin和用户社交信息UserInfo
func UserRegister(login *model.UserLogin, info *model.UserInfo) error {
	return global.DOUYIN_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&login).Error; err != nil {
			return err
		}
		if err := tx.Create(&info).Error; err != nil {
			return err
		}
		return nil
	})
}
