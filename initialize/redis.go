package initialize

import (
	"context"
	"fmt"
	"github.com/linzijie1998/bytedance_camp_douyin/global"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func InitRedis() *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:     global.DOUYIN_CONFIG.Redis.Addr(),
		Password: global.DOUYIN_CONFIG.Redis.Password,
		DB:       global.DOUYIN_CONFIG.Redis.DB,
	})
	if _, err := r.Ping(ctx).Result(); err != nil {
		global.DOUYIN_LOGGER.Debug(fmt.Sprintf(
			"Redis服务连接失败 Addr: %s, Password: %s, DB: %d",
			global.DOUYIN_CONFIG.Redis.Addr(),
			global.DOUYIN_CONFIG.Redis.Password,
			global.DOUYIN_CONFIG.Redis.DB))
		panic(err)
	}
	global.DOUYIN_LOGGER.Info(fmt.Sprintf(
		"Redis服务连接成功 Addr: %s, Password: %s, DB: %d",
		global.DOUYIN_CONFIG.Redis.Addr(),
		global.DOUYIN_CONFIG.Redis.Password,
		global.DOUYIN_CONFIG.Redis.DB))
	return r
}
