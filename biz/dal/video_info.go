package dal

import (
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"github.com/linzijie1998/bytedance_camp_douyin/model"
	"time"
)

// CreateVideoInfo 添加视频信息
func CreateVideoInfo(info *model.Video) error {
	return global.DOUYIN_DB.Create(info).Error
}

// DeleteVideoInfoByID 根据视频ID删除视频信息
func DeleteVideoInfoByID(id int64) error {
	return global.DOUYIN_DB.Where("id = ?", id).Delete(&model.Video{}).Error
}

// UpdateVideoInfo 更新视频信息
func UpdateVideoInfo(info *model.Video) error {
	// 此处应该使用Save方法而不是Updates方法
	//return global.DOUYIN_DB.Updates(info).Error // 0值不会存储
	return global.DOUYIN_DB.Save(info).Error
}

// QueryVideoInfoByUserInfoID 根据UserInfoID查询该用户所有发布的视频
func QueryVideoInfoByUserInfoID(userInfoID int64) (infos []model.Video, err error) {
	err = global.DOUYIN_DB.Find(&infos, "user_info_id = ?", userInfoID).Error
	return
}

// QueryVideoInfoByID 根据ID查询视频信息
func QueryVideoInfoByID(id int64) (infos []model.Video, err error) {
	err = global.DOUYIN_DB.Find(&infos, "id = ?", id).Error
	return
}

// QueryVideoInfosWithLimit 根据发布时间顺序查询视频（无时间限制）
func QueryVideoInfosWithLimit(limit int) (infos []model.Video, err error) {
	err = global.DOUYIN_DB.Order("created_at ASC").Limit(limit).Find(&infos).Error
	return
}

// QueryVideoInfosWithLimitAndTime 根据发布时间顺序查询视频（有时间限制）
func QueryVideoInfosWithLimitAndTime(limit int, latestTime time.Time) (infos []model.Video, err error) {
	err = global.DOUYIN_DB.Where("created_at < ?", latestTime).Order("created_at ASC").Limit(limit).Find(&infos).Error
	return
}
