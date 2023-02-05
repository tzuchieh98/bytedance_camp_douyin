package dal

import (
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"github.com/linzijie1998/bytedance_camp_douyin/model"
)

func CreateLoginInfo(login *model.UserLogin) error {
	return global.DOUYIN_DB.Create(login).Error
}

func DeleteLoginInfo(id int64) error {
	return global.DOUYIN_DB.Where("id = ?", id).Delete(&model.UserLogin{}).Error
}

func UpdateLoginInfo(login *model.UserLogin) error {
	return global.DOUYIN_DB.Updates(login).Error
}

func QueryLoginInfoByUsername(username string) (loginInfos []model.UserLogin, err error) {
	err = global.DOUYIN_DB.Find(&loginInfos, "username = ?", username).Error
	return
}
