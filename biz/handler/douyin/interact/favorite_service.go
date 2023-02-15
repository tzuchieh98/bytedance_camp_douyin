// Code generated by hertz generator.

package interact

import (
	"context"
	"fmt"
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"path"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/linzijie1998/bytedance_camp_douyin/biz/cache"
	"github.com/linzijie1998/bytedance_camp_douyin/biz/dal"
	"github.com/linzijie1998/bytedance_camp_douyin/biz/model/douyin/base"
	interact "github.com/linzijie1998/bytedance_camp_douyin/biz/model/douyin/interact"
	"github.com/linzijie1998/bytedance_camp_douyin/util"
)

const (
	favoriteActionLike   = 1
	favoriteActionCancel = 2
)

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	// 1. 从token中解析user_id
	// 2. 查询该用户的点赞状态
	// 3. 根据action_type进行处理
	var err error
	var req interact.FavoriteActionReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interact.FavoriteActionResp)

	rawID, exists := c.Get("token_user_id")
	if !exists {
		global.DOUYIN_LOGGER.Debug("未从上下文中解析到USERID")
		resp.StatusCode = 1
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	userID := int64(rawID.(uint))
	isFavorite, err := cache.GetFavoriteState(userID, req.VideoID)
	if err != nil {
		global.DOUYIN_LOGGER.Debug(fmt.Sprintf("点赞数据查询失败 err: %v", err))
		resp.StatusCode = 1
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	// 1. 判断是否已经点赞
	// 2. 更新点赞状态
	// 3. 添加点赞计数
	if req.ActionType == favoriteActionLike {
		// 判断是否已经点赞
		if isFavorite {
			global.DOUYIN_LOGGER.Info(fmt.Sprintf("ID为%d的用户尝试重复对ID为%d视频点赞", userID, req.VideoID))
			resp.StatusCode = 1
			c.JSON(consts.StatusBadRequest, resp)
			return
		}
		// 更新点赞状态
		if err := cache.UpdateFavoriteState(userID, req.VideoID, true); err != nil {
			global.DOUYIN_LOGGER.Debug(fmt.Sprintf("视频点赞状态更新失败 err: %v", err))
			resp.StatusCode = 1
			c.JSON(consts.StatusInternalServerError, resp)
			return
		}
		// 更新点赞计数
		if err := cache.UpdateFavoriteCount(req.VideoID, true); err != nil {
			global.DOUYIN_LOGGER.Debug(fmt.Sprintf("视频点赞计数更新失败 err: %v", err))
			resp.StatusCode = 1
			c.JSON(consts.StatusInternalServerError, resp)
			return
		}

	} else if req.ActionType == favoriteActionCancel {
		// 判断是否已经点赞
		if !isFavorite {
			resp.StatusCode = 1
			global.DOUYIN_LOGGER.Info(fmt.Sprintf("ID为%d的用户尝试对未点赞的ID为%d视频取消点赞", userID, req.VideoID))
			c.JSON(consts.StatusBadRequest, resp)
			return
		}
		// 更新点赞状态
		if err := cache.UpdateFavoriteState(userID, req.VideoID, false); err != nil {
			global.DOUYIN_LOGGER.Debug(fmt.Sprintf("视频点赞信息更新失败 err: %v", err))
			resp.StatusCode = 1
			c.JSON(consts.StatusInternalServerError, resp)
			return
		}
		// 更新点赞计数
		if err := cache.UpdateFavoriteCount(req.VideoID, false); err != nil {
			global.DOUYIN_LOGGER.Debug(fmt.Sprintf("视频点赞计数更新失败 err: %v", err))
			resp.StatusCode = 1
			c.JSON(consts.StatusInternalServerError, resp)
			return
		}
	} else {
		global.DOUYIN_LOGGER.Info(fmt.Sprintf("错误的点赞操作 action_type: %d", req.ActionType))
		resp.StatusCode = 1
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req interact.FavoriteListReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(interact.FavoriteListResp)

	if req.Token == "" {
		c.JSON(consts.StatusOK, resp)
		return
	}

	j := util.NewJWT()
	claim, err := j.ParseToken(req.Token)
	if err != nil {
		global.DOUYIN_LOGGER.Info(fmt.Sprintf("Token解析失败 err: %v", err))
		resp.StatusCode = 1
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	userID := int64(claim.UserInfo.ID)

	videoIDs, err := cache.QueryFavoriteVideosByUserID(userID)
	if err != nil {
		global.DOUYIN_LOGGER.Debug(fmt.Sprintf("查询点赞视频失败 err: %v", err))
		resp.StatusCode = 1
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	videoList := make([]*base.Video, len(videoIDs))

	for i, videoId := range videoIDs {
		videoInfos, err := dal.QueryVideoInfoByID(videoId)
		if err != nil {
			global.DOUYIN_LOGGER.Debug(fmt.Sprintf("查询视频信息失败 err: %v", err))
			resp.StatusCode = 1
			c.JSON(consts.StatusInternalServerError, resp)
			return
		}
		if len(videoInfos) != 1 {
			global.DOUYIN_LOGGER.Warn(fmt.Sprintf("查询到%d条的ID为%d的视频信息", len(videoInfos), videoId))
			resp.StatusCode = 1
			c.JSON(consts.StatusInternalServerError, resp)
			return
		}

		userInfos, err := dal.QueryUserInfoByUserID(videoInfos[0].UserInfoID)
		if err != nil {
			global.DOUYIN_LOGGER.Debug(fmt.Sprintf("用户信息查询失败 err: %v", err))
			resp.StatusCode = 1
			c.JSON(consts.StatusInternalServerError, resp)
			return
		}
		if len(userInfos) != 1 {
			global.DOUYIN_LOGGER.Warn(fmt.Sprintf("查询到超过一条的ID为%d的用户信息", videoInfos[0].UserInfoID))
			resp.StatusCode = 1
			c.JSON(consts.StatusInternalServerError, resp)
			return
		}

		// 从Redis中查询点赞计数和评论计数
		favoriteCnt, _ := cache.GetFavoriteCount(int64(userInfos[0].ID))
		commentCnt, _ := cache.GetCommentCount(int64(userInfos[0].ID))
		followCnt, _ := cache.GetFollowCount(int64(videoInfos[0].ID))
		followerCnt, _ := cache.GetFollowerCount(int64(videoInfos[0].ID))

		var user base.User
		user.ID = int64(userInfos[0].ID)
		user.Name = userInfos[0].Name
		user.FollowCount = &followCnt
		user.FollowerCount = &followerCnt

		var video base.Video
		video.ID = int64(videoInfos[0].ID)
		video.Author = &user
		video.CommentCount = commentCnt
		video.FavoriteCount = favoriteCnt
		video.PlayURL = util.GetPlayURLByFilename(path.Base(videoInfos[0].VideoPath))
		video.CoverURL = util.GetCoverURLByFilename(path.Base(videoInfos[0].CoverPath))
		video.IsFavorite = true // 点赞列表
		video.Title = videoInfos[0].Title

		videoList[i] = &video
	}

	resp.VideoList = videoList
	c.JSON(consts.StatusOK, resp)
}
