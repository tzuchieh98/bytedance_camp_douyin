package cache

import (
	"context"
	"fmt"
	"strconv"

	"github.com/linzijie1998/bytedance_camp_douyin/global"
)

const (
	FollowKey   = "follow_uid_%d"
	FollowerKey = "follower_uid_%d"
)

var ctx = context.Background()

// UpdateFollowAndFollowerState 更新用户的关注状态
func UpdateFollowAndFollowerState(userID, targetID int64, isAdd bool) (err error) {
	followKey := fmt.Sprintf(FollowKey, userID)
	followerKey := fmt.Sprintf(FollowerKey, targetID)
	if isAdd {
		if err = global.DOUYIN_REDIS.SAdd(ctx, followKey, targetID).Err(); err != nil {
			return
		}
		if err = global.DOUYIN_REDIS.SAdd(ctx, followerKey, userID).Err(); err != nil {
			return
		}
	} else {
		if err = global.DOUYIN_REDIS.SRem(ctx, followKey, targetID).Err(); err != nil {
			return
		}
		if err = global.DOUYIN_REDIS.SRem(ctx, followerKey, userID).Err(); err != nil {
			return
		}
	}
	return
}

// GetFollowerCount 获取用户的粉丝数
func GetFollowerCount(userID int64) (int64, error) {
	key := fmt.Sprintf(FollowerKey, userID)
	return global.DOUYIN_REDIS.SCard(ctx, key).Result()
}

// GetFollowCount 获取用户的关注数
func GetFollowCount(userID int64) (int64, error) {
	key := fmt.Sprintf(FollowKey, userID)
	return global.DOUYIN_REDIS.SCard(ctx, key).Result()
}

// GetFollowState 获取用户是否关注
func GetFollowState(userID int64, followerID int64) (bool, error) {
	key := fmt.Sprintf(FollowKey, userID)
	return global.DOUYIN_REDIS.SIsMember(ctx, key, followerID).Result()
}

// QueryFollowerByUserID 获取粉丝列表
func QueryFollowerByUserID(userID int64) (userList []int64, err error) {
	key := fmt.Sprintf(FollowerKey, userID)
	val, err := global.DOUYIN_REDIS.SMembers(ctx, key).Result()
	if err != nil {
		return []int64{}, err
	}

	userList = make([]int64, len(val))

	for i, id := range val {
		user, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return []int64{}, err
		}
		userList[i] = user
	}
	return userList, nil
}

// QueryFollowByUserID 获取关注列表
func QueryFollowByUserID(userID int64) (userList []int64, err error) {
	key := fmt.Sprintf(FollowKey, userID)
	val, err := global.DOUYIN_REDIS.SMembers(ctx, key).Result()
	if err != nil {
		return []int64{}, err
	}

	userList = make([]int64, len(val))

	for i, id := range val {
		user, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return []int64{}, err
		}
		userList[i] = user
	}
	return userList, nil
}
