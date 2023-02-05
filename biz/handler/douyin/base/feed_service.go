// Code generated by hertz generator.

package base

import (
	"context"
	"fmt"
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"github.com/linzijie1998/bytedance_camp_douyin/model"
	"path"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/linzijie1998/bytedance_camp_douyin/biz/cache"
	"github.com/linzijie1998/bytedance_camp_douyin/biz/dal"
	base "github.com/linzijie1998/bytedance_camp_douyin/biz/model/douyin/base"
	"github.com/linzijie1998/bytedance_camp_douyin/util"
)

const (
	MaxVideosNum = 30
)

// Feed 视频流信息.
// @router /douyin/feed/ [GET]
func Feed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req base.FeedReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp := new(base.FeedResp)

	var userID int64

	// 通过token是否为空判断登录状态, 如果登录则解析token
	if req.Token != nil {
		j := util.NewJWT()
		claims, err := j.ParseToken(*req.Token)
		if err != nil {
			global.DOUYIN_LOGGER.Debug("Token解析错误")
			resp.StatusCode = 1
			return
		}
		userID = int64(claims.UserInfo.ID)
	}

	// 获取视频流信息
	var videoInfos []model.Video
	// LatestTime为空则将其设置为当前时间
	if req.LatestTime == nil {
		*req.LatestTime = time.Now().UnixNano() / 1e6
	}
	videoInfos, err = dal.QueryVideoInfosWithLimitAndTime(
		MaxVideosNum, time.Unix(*req.LatestTime/1e3, *req.LatestTime/1e3))
	if err != nil {
		global.DOUYIN_LOGGER.Debug(fmt.Sprintf("查询视频信息失败 err:%v", err))
		resp.StatusCode = 1
		c.JSON(consts.StatusInternalServerError, resp)
	}

	// 处理返回结构体
	videoList := make([]*base.Video, len(videoInfos))
	for i, info := range videoInfos {
		// 查询发布者信息
		userInfos, err := dal.QueryUserInfoByUserID(info.UserInfoID)
		if err != nil || len(userInfos) != 1 {
			global.DOUYIN_LOGGER.Debug(fmt.Sprintf("查询用户信息失败 err:%v", err))
			resp.StatusCode = 1
			c.JSON(consts.StatusInternalServerError, resp)
			return
		}
		// 查询点赞状态
		isFavorite := false
		if req.Token != nil {
			// 登录状态查询点赞状态
			isFavorite, err = cache.GetFavoriteState(userID, int64(info.ID))
			if err != nil {
				global.DOUYIN_LOGGER.Debug(fmt.Sprintf("查询点赞状态失败 err:%v", err))
				resp.StatusCode = 1
				c.JSON(consts.StatusInternalServerError, resp)
				return
			}
		}

		var user = new(base.User)
		user.ID = int64(userInfos[0].ID)
		user.Name = userInfos[0].Name
		user.FollowCount = &userInfos[0].FollowCount
		user.FollowerCount = &userInfos[0].FollowerCount
		user.IsFollow = true

		var video = new(base.Video)
		video.ID = int64(info.ID)
		video.PlayURL = util.GetPlayURLByFilename(path.Base(info.VideoPath))
		video.CoverURL = util.GetCoverURLByFilename(path.Base(info.CoverPath))
		video.FavoriteCount = info.FavoriteCount
		video.CommentCount = info.CommentCount
		video.Author = user
		video.IsFavorite = isFavorite
		videoList[i] = video
	}

	nextTime := time.Now().UnixNano() / 1e6
	resp.VideoList = videoList
	resp.NextTime = &nextTime
	c.JSON(consts.StatusOK, resp)
}