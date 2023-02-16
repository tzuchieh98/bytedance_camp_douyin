package cache

import (
	"fmt"
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"strconv"
)

const commentCountKey = "comment_cnt_vid_%d"

// UpdateCommentCount 更新评论计数, isAdd为true则点赞数加1, 为false则点赞数减1, 不做计数的正负判断
func UpdateCommentCount(videoID int64, isAdd bool) (err error) {
	key := fmt.Sprintf(commentCountKey, videoID)
	if isAdd {
		err = global.DOUYIN_REDIS.Incr(ctx, key).Err()
	} else {
		err = global.DOUYIN_REDIS.Decr(ctx, key).Err()
	}
	return
}

// GetCommentCount 获取视频点赞计数, 不做计数的正负判断
func GetCommentCount(videoID int64) (int64, error) {
	key := fmt.Sprintf(commentCountKey, videoID)
	commentCnt, err := global.DOUYIN_REDIS.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(commentCnt, 10, 64)
}
