package cache

import (
	"fmt"
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"strconv"
)

const (
	FavoriteVideoKey    = "favorite_set_uid_%d"
	UserKey             = "user_%d"
	VideoFavoriteField  = "video_favorite_cnt"
	UserFavoriteField   = "user_favorite_cnt"
	TotalFavoritedField = "total_favorited_cnt"
)

// UpdateFavoriteState 更新用户的视频点赞状态 视频点赞计数 和 作者获赞计数
func UpdateFavoriteState(userID, videoID, authID int64, state bool) (err error) {
	key := fmt.Sprintf(FavoriteVideoKey, userID)
	videoKey := fmt.Sprintf(VideoKey, videoID)
	userKey := fmt.Sprintf(UserKey, userID)
	authKey := fmt.Sprintf(UserKey, authID)

	var is int64 = 0
	if state {
		// 点赞
		if is, err = global.DOUYIN_REDIS.SAdd(ctx, key, videoID).Result(); is == 0 {
			return
		}
		err = global.DOUYIN_REDIS.HIncrBy(ctx, videoKey, VideoFavoriteField, 1).Err()
		err = global.DOUYIN_REDIS.HIncrBy(ctx, userKey, UserFavoriteField, 1).Err()
		err = global.DOUYIN_REDIS.HIncrBy(ctx, authKey, TotalFavoritedField, 1).Err()
	} else {
		// 取消点赞
		is, err = global.DOUYIN_REDIS.SRem(ctx, key, videoID).Result()
		err = global.DOUYIN_REDIS.HIncrBy(ctx, videoKey, VideoFavoriteField, -1).Err()
		err = global.DOUYIN_REDIS.HIncrBy(ctx, userKey, UserFavoriteField, -1).Err()
		err = global.DOUYIN_REDIS.HIncrBy(ctx, authKey, TotalFavoritedField, -1).Err()
	}
	return
}

// GetFavoriteState 查询用户的视频点赞状态
func GetFavoriteState(userID int64, videoID int64) (bool, error) {
	key := fmt.Sprintf(FavoriteVideoKey, userID)
	return global.DOUYIN_REDIS.SIsMember(ctx, key, videoID).Result()
}

// QueryFavoriteVideosByUserID 查询某个用户的所有点赞视频ID
func QueryFavoriteVideosByUserID(userID int64) (videoList []int64, err error) {
	key := fmt.Sprintf(FavoriteVideoKey, userID)
	val, err := global.DOUYIN_REDIS.SMembers(ctx, key).Result()
	if err != nil {
		return []int64{}, err
	}

	videoList = make([]int64, len(val))

	for i, id := range val {
		videoID, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return []int64{}, err
		}
		videoList[i] = videoID
	}
	return videoList, nil
}

//// UpdateFavoriteCount 更新视频点赞计数, isAdd为true则点赞数加1, 为false则点赞数减1, 不做计数的正负判断
//func UpdateFavoriteCount(videoID int64, isAdd bool) (err error) {
//	key := fmt.Sprintf(favoriteCountKey, videoID)
//	if isAdd {
//		err = global.DOUYIN_REDIS.Incr(ctx, key).Err()
//	} else {
//		err = global.DOUYIN_REDIS.Decr(ctx, key).Err()
//	}
//	return
//}
//
//// GetFavoriteCount 获取视频点赞计数, 不做计数的正负判断
//func GetFavoriteCount(videoID int64) (int64, error) {
//	key := fmt.Sprintf(favoriteCountKey, videoID)
//	favoriteCnt, err := global.DOUYIN_REDIS.Get(ctx, key).Result()
//	if err != nil {
//		return 0, err
//	}
//	return strconv.ParseInt(favoriteCnt, 10, 64)
//}
