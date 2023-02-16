package cache

import (
	"context"
	"fmt"
	"strconv"

	"github.com/linzijie1998/bytedance_camp_douyin/global"
)

const (
	followCountKey   = "follow_cnt_uid_%d"
	followerCountKey = "follower_cnt_uid_%d"
	followUserKey    = "follow_user_uid_%d"
	followerUserKey  = "follower_user_uid_%d"
)

var ctx = context.Background()

func UpdateFollowerCount(userID int64, isAdd bool) (err error) {
	key := fmt.Sprintf(followerCountKey, userID)
	if isAdd {
		err = global.DOUYIN_REDIS.Incr(ctx, key).Err()
	} else {
		err = global.DOUYIN_REDIS.Decr(ctx, key).Err()
	}
	return
}

func GetFollowerCount(userID int64) (int64, error) {
	key := fmt.Sprintf(followerCountKey, userID)
	followerCnt, err := global.DOUYIN_REDIS.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(followerCnt, 10, 64)
}

func UpdateFollowCount(userID int64, isAdd bool) (err error) {
	key := fmt.Sprintf(followCountKey, userID)
	if isAdd {
		err = global.DOUYIN_REDIS.Incr(ctx, key).Err()
	} else {
		err = global.DOUYIN_REDIS.Decr(ctx, key).Err()
	}
	return
}

func GetFollowCount(userID int64) (int64, error) {
	key := fmt.Sprintf(followCountKey, userID)
	followCnt, err := global.DOUYIN_REDIS.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(followCnt, 10, 64)
}

// UpdateFollowerState .
func UpdateFollowerState(userID int64, followerID int64, state bool) (err error) {
	key := fmt.Sprintf(followerUserKey, userID)
	if state {
		_, err = global.DOUYIN_REDIS.SAdd(ctx, key, followerID).Result()
	} else {
		_, err = global.DOUYIN_REDIS.SRem(ctx, key, followerID).Result()
	}
	return
}

// GetFollowerState .
func GetFollowerState(userID int64, followerID int64) (bool, error) {
	key := fmt.Sprintf(followerUserKey, followerID)
	return global.DOUYIN_REDIS.SIsMember(ctx, key, followerID).Result()
}

func QueryFollowerByUserID(userID int64) ([]int64, error) {
	key := fmt.Sprintf(followerUserKey, userID)
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
	key := fmt.Sprintf(followUserKey, userID)
	if state {
		_, err = global.DOUYIN_REDIS.SAdd(ctx, key, followID).Result()
	} else {
		_, err = global.DOUYIN_REDIS.SRem(ctx, key, followID).Result()
	}
	return
}

func GetFollowState(userID int64, followID int64) (bool, error) {
	key := fmt.Sprintf(followUserKey, userID)
	return global.DOUYIN_REDIS.SIsMember(ctx, key, followID).Result()
}

func QueryFollowByUserID(userID int64) ([]int64, error) {
	key := fmt.Sprintf(followUserKey, userID)
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
