package dal

import (
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"github.com/linzijie1998/bytedance_camp_douyin/model"
	"gorm.io/gorm"
)

// CreateComment 添加评论
func CreateComment(comment *model.Comment) error {
	return global.DOUYIN_DB.Create(comment).Error
}

// DeleteCommentByID 删除评论
func DeleteCommentByID(id int64) error {
	return global.DOUYIN_DB.Where("id = ?", id).Delete(&model.Comment{}).Error
}

// QueryCommentByVideoID 根据视频ID倒序查询该视频的所有评论
func QueryCommentByVideoID(videoId int64) (comments []model.Comment, err error) {
	err = global.DOUYIN_DB.Find(&comments, "video_id = ?", videoId).Order("created_at desc").Error
	return
}

// PublishComment 同时更新视频信息和评论信息
func PublishComment(comment *model.Comment, video *model.Video) error {
	return global.DOUYIN_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&comment).Error; err != nil {
			return err
		}
		if err := tx.Updates(&video).Error; err != nil {
			return err
		}
		return nil
	})
}
