package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	UserInfoID int64  `json:"user_info_id"` // 上传者的用户信息id
	Title      string `json:"title"`        // 视频标题
	VideoPath  string `json:"play_url"`     // 视频存储位置
	CoverPath  string `json:"cover_url"`    // 封面存储位置
	//FavoriteCount int64  `json:"favorite_count"` // 点赞数
	//CommentCount  int64  `json:"comment_count"`  // 评论数
}

func (Video) TableName() string {
	return "videos"
}
