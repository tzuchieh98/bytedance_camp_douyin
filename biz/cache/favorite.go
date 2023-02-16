package cache

import (
	"fmt"
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"strconv"
)

const (
	favoriteVideoKey = "favorite_list_uid_%d"
	favoriteCountKey = "favorite_cnt_vid_%d"
)

// UpdateFavoriteState 更新用户的视频点赞状态
func UpdateFavoriteState(userID int64, videoID int64, state bool) (err error) {
	key := fmt.Sprintf(favoriteVideoKey, userID)
	if state {
		// 点赞
		err = global.DOUYIN_REDIS.SAdd(ctx, key, videoID).Err()
	} else {
		// 取消点赞
		_, err = global.DOUYIN_REDIS.SRem(ctx, key, videoID).Result()
	}
	return
}

// GetFavoriteState 查询用户的视频点赞状态
func GetFavoriteState(userID int64, videoID int64) (bool, error) {
	key := fmt.Sprintf(favoriteVideoKey, userID)
	return global.DOUYIN_REDIS.SIsMember(ctx, key, videoID).Result()
}

// QueryFavoriteVideosByUserID 查询某个用户的所有点赞视频ID
func QueryFavoriteVideosByUserID(userID int64) ([]int64, error) {
	key := fmt.Sprintf(favoriteVideoKey, userID)
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

// UpdateFavoriteCount 更新视频点赞计数, isAdd为true则点赞数加1, 为false则点赞数减1, 不做计数的正负判断
func UpdateFavoriteCount(videoID int64, isAdd bool) (err error) {
	key := fmt.Sprintf(favoriteCountKey, videoID)
	if isAdd {
		err = global.DOUYIN_REDIS.Incr(ctx, key).Err()
	} else {
		err = global.DOUYIN_REDIS.Decr(ctx, key).Err()
	}
	return
}

// GetFavoriteCount 获取视频点赞计数, 不做计数的正负判断
func GetFavoriteCount(videoID int64) (int64, error) {
	key := fmt.Sprintf(favoriteCountKey, videoID)
	favoriteCnt, err := global.DOUYIN_REDIS.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(favoriteCnt, 10, 64)
}
