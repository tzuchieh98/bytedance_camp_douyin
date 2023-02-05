package cache

import (
	"context"
	"fmt"
	"strconv"

	"github.com/linzijie1998/bytedance_camp_douyin/global"
)

const (
	FAVORITE = "favorite"
	FOLLOW   = "follow"
	FOLLOWER = "follower"
)

var ctx = context.Background()

// UpdateFavoriteState 更新视频点赞状态
func UpdateFavoriteState(userID int64, videoID int64, state bool) (err error) {
	key := fmt.Sprintf("%s:%d", FAVORITE, userID)
	if state {
		// 点赞
		_, err = global.DOUYIN_REDIS.SAdd(ctx, key, videoID).Result()
	} else {
		// 取消点赞
		_, err = global.DOUYIN_REDIS.SRem(ctx, key, videoID).Result()
	}
	return
}

// GetFavoriteState 查询视频点赞状态
func GetFavoriteState(userID int64, videoID int64) (bool, error) {
	key := fmt.Sprintf("%s:%d", FAVORITE, userID)
	return global.DOUYIN_REDIS.SIsMember(ctx, key, videoID).Result()
}

// QueryFavoriteVideosByUserID 查询某个用户的所有点赞视频ID
func QueryFavoriteVideosByUserID(userID int64) ([]int64, error) {
	key := fmt.Sprintf("%s:%d", FAVORITE, userID)
	ssc := global.DOUYIN_REDIS.SMembers(ctx, key)
	_, err := ssc.Result()
	if err != nil {
		return []int64{}, err
	}
	val := ssc.Val()
	videoList := make([]int64, len(val))
	for i, id := range val {
		videoID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return []int64{}, err
		}
		videoList[i] = videoID
	}
	return videoList, nil
}

// UpdateFollowerState .
func UpdateFollowerState(userID int64, followerID int64, state bool) (err error) {
	key := fmt.Sprintf("%s:%d", FOLLOWER, userID)
	if state {
		_, err = global.DOUYIN_REDIS.SAdd(ctx, key, followerID).Result()
	} else {
		_, err = global.DOUYIN_REDIS.SRem(ctx, key, followerID).Result()
	}
	return
}

// GetFollowerState .
func GetFollowerState(userID int64, followerID int64) (bool, error) {
	key := fmt.Sprintf("%s:%d", FOLLOWER, followerID)
	return global.DOUYIN_REDIS.SIsMember(ctx, key, followerID).Result()
}

func QueryFollowerByUserID(userID int64) ([]int64, error) {
	key := fmt.Sprintf("%s:%d", FOLLOWER, userID)
	ssc := global.DOUYIN_REDIS.SMembers(ctx, key)
	_, err := ssc.Result()
	if err != nil {
		return []int64{}, err
	}
	val := ssc.Val()
	videoList := make([]int64, len(val))
	for i, id := range val {
		videoID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return []int64{}, err
		}
		videoList[i] = videoID
	}
	return videoList, nil
}

// UpdateFollowState .
func UpdateFollowState(userID int64, followID int64, state bool) (err error) {
	key := fmt.Sprintf("%s:%d", FOLLOW, userID)
	if state {
		_, err = global.DOUYIN_REDIS.SAdd(ctx, key, followID).Result()
	} else {
		_, err = global.DOUYIN_REDIS.SRem(ctx, key, followID).Result()
	}
	return
}

func GetFollowState(userID int64, followID int64) (bool, error) {
	key := fmt.Sprintf("%s:%d", FOLLOW, followID)
	return global.DOUYIN_REDIS.SIsMember(ctx, key, followID).Result()
}

func QueryFollowByUserID(userID int64) ([]int64, error) {
	key := fmt.Sprintf("%s:%d", FOLLOW, userID)
	ssc := global.DOUYIN_REDIS.SMembers(ctx, key)
	_, err := ssc.Result()
	if err != nil {
		return []int64{}, err
	}
	val := ssc.Val()
	videoList := make([]int64, len(val))
	for i, id := range val {
		videoID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return []int64{}, err
		}
		videoList[i] = videoID
	}
	return videoList, nil
}
