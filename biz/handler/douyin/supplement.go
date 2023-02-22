package douyin

import (
	"errors"
	"fmt"
	"github.com/linzijie1998/bytedance_camp_douyin/biz/cache"
	"github.com/linzijie1998/bytedance_camp_douyin/biz/dal"
	"github.com/linzijie1998/bytedance_camp_douyin/biz/model/douyin/base"
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"github.com/linzijie1998/bytedance_camp_douyin/model"
	"github.com/linzijie1998/bytedance_camp_douyin/util"
	"path"
)

func VideoInfoSupplement(userID int64, video *base.Video, info *model.Video) (err error) {

	// 视频信息
	if info == nil {
		var videos []model.Video
		videos, err = dal.QueryVideoInfoByID(video.ID)
		if err != nil && len(videos) != 1 {
			global.DOUYIN_LOGGER.Debug(fmt.Sprintf("查询视频信息失败 err: %v", err))
			return
		}
		info = &videos[0]
		if info == nil {
			return errors.New("video is not exist")
		}
	}
	video.Title = info.Title
	video.PlayURL = util.GetPlayURLByFilename(path.Base(info.VideoPath))
	video.CoverURL = util.GetCoverURLByFilename(path.Base(info.CoverPath))

	// 从Redis中查询点赞计数和评论计数
	videoCount, err := cache.GetVideoCount(video.ID)
	if err != nil {
		fmt.Println(2, err)
		return
	}
	video.FavoriteCount = videoCount[cache.VideoFavoriteField]
	video.CommentCount = videoCount[cache.CommentField]

	// 查找是否点赞
	var isFavorite = false
	if userID != 0 {
		// 登录状态查询点赞状态
		isFavorite, err = cache.GetFavoriteState(userID, video.ID)
		if err != nil {
			global.DOUYIN_LOGGER.Debug(fmt.Sprintf("查询点赞状态失败 err:%v", err))
			return
		}
	}
	video.IsFavorite = isFavorite

	return
}

func UserInfoSupplement(userID int64, user *base.User, info *model.UserInfo) (err error) {

	// 用户信息
	if info == nil {
		var userInfos = make([]model.UserInfo, 1)
		userInfos, err = dal.QueryUserInfoByUserID(user.ID)
		if err != nil || len(userInfos) != 1 {
			global.DOUYIN_LOGGER.Debug(fmt.Sprintf("用户信息查询失败 err: %v", err))
			return
		}
		info = &userInfos[0]
		if info == nil {
			return errors.New("user is not exist")
		}
	}
	user.Name = info.Name
	//user.Avatar = &model.Avatar
	//user.Signature = &model.Signature
	//user.BackgroundImage = &model.BackgroundImage

	// 关注数据
	followCount, err := cache.GetFollowCount(user.ID)
	followerCount, err := cache.GetFollowerCount(user.ID)
	if err != nil {
		fmt.Println(6, err)
		return
	}
	user.FollowCount = &followCount
	user.FollowerCount = &followerCount

	// 点赞数
	userCount, err := cache.GetUserCount(user.ID)
	if err != nil {
		return
	}
	user.FavoriteCount = new(int64)
	*user.FavoriteCount = userCount[cache.UserFavoriteField]
	user.TotalFavorited = new(int64)
	*user.TotalFavorited = userCount[cache.TotalFavoritedField]
	user.WorkCount = new(int64)
	*user.WorkCount = userCount[cache.WorkField]

	// 查找是否关注
	var isFollow = false
	if userID != 0 && userID != user.ID {
		isFollow, err = cache.GetFollowState(userID, user.ID)
		if err != nil {
			fmt.Println(9, err)
			return
		}
	}
	user.IsFollow = isFollow

	return
}
