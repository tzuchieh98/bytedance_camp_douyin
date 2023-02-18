package dal

import (
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"github.com/linzijie1998/bytedance_camp_douyin/model"
)

func CreateMessage(msg *model.Message) error {
	return global.DOUYIN_DB.Create(msg).Error
}

func QueryMessageByUserIDAndToUserID(userID int64, toUserID int64) (messages []model.Message, err error) {
	err = global.DOUYIN_DB.Find(&messages, "user_id = ? and to_user_id = ?", userID, toUserID).Error
	return
}

func QueryMessageByUserIDAndToUserIDWithLimit(userID int64, toUserID int64, limit int64) (messages []model.Message, err error) {
	err = global.DOUYIN_DB.Where("publish_date > ?", limit).Find(&messages, "user_id = ? and to_user_id = ?", userID, toUserID).Error
	return
}
