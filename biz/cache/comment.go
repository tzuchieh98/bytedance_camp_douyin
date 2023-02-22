package cache

import (
	"fmt"
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"strconv"
)

const (
	VideoKey     = "video_%d"
	CommentField = "comment_cnt"
	WorkField    = "work_cnt"
)

// GetUserCount 获取用户计数 作品数 获赞数 点赞数
func GetUserCount(userID int64) (countMap map[string]int64, err error) {
	key := fmt.Sprintf(UserKey, userID)
	val, err := global.DOUYIN_REDIS.HGetAll(ctx, key).Result()
	if err != nil {
		return
	}

	countMap = make(map[string]int64)

	for k, v := range val {
		if countMap[k], err = strconv.ParseInt(v, 10, 64); err != nil {
			return
		}
	}

	return
}

// UpdateWorkCount 更新作品数
func UpdateWorkCount(userID int64) error {
	key := fmt.Sprintf(UserKey, userID)
	return global.DOUYIN_REDIS.HIncrBy(ctx, key, WorkField, 1).Err()
}

// UpdateCommentCount 更新评论计数, isAdd为true则点赞数加1, 为false则点赞数减1, 不做计数的正负判断
func UpdateCommentCount(videoID int64, isAdd bool) (err error) {
	key := fmt.Sprintf(VideoKey, videoID)
	if isAdd {
		err = global.DOUYIN_REDIS.HIncrBy(ctx, key, CommentField, 1).Err()
	} else {
		err = global.DOUYIN_REDIS.HIncrBy(ctx, key, CommentField, -1).Err()
	}
	return
}

// GetVideoCount 获取视频点赞 评论数, 不做计数的正负判断
func GetVideoCount(videoID int64) (countMap map[string]int64, err error) {
	key := fmt.Sprintf(VideoKey, videoID)
	val, err := global.DOUYIN_REDIS.HGetAll(ctx, key).Result()
	if err != nil {
		return
	}

	countMap = make(map[string]int64)

	for k, v := range val {
		if countMap[k], err = strconv.ParseInt(v, 10, 64); err != nil {
			return
		}
	}

	return
}

// GetCommentCount 获取视频点赞计数, 不做计数的正负判断
//func GetCommentCount(videoID int64) (int64, error) {
//	key := fmt.Sprintf(commentCountKey, videoID)
//	commentCnt, err := global.DOUYIN_REDIS.Get(ctx, key).Result()
//	if err != nil {
//		return 0, err
//	}
//	return strconv.ParseInt(commentCnt, 10, 64)
//}
